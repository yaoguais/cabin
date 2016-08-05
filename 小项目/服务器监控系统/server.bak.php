<?php
// 打开连接时,推送10个数据过去
// 然后新开一个task，去发送消息
// 采取一个策略, 每次采集的数据都保存到文件(data/machine-YYYYMMdd.log)中去，保存的格式是:
// 时间戳 CPU 内存

class G{
    public static $clients = null;
    public static $sshConfig = null;
}

G::$sshConfig = require __DIR__ .'/config.php';

$server = new swoole_websocket_server("0.0.0.0", 9505);
$server->set(['worker_num' => 1, 'daemonize' => 1, 'log_file' => __DIR__ . '/data/log_file.log']);

$server->on('open', function (swoole_websocket_server $server, $request){

    if( !G::$clients ){
        $server->tick(3000, function () use ($server){

            if( G::$clients !== null ){
                foreach( G::$clients as $fd ){
                    $server->push($fd, get_data(1));
                }
            }
        });
    }
    G::$clients[$request->fd] = $request->fd;
    $info = $server->connection_info($request->fd);
    echo "server: handshake success with fd{$request->fd} ", (isset($info['remote_ip']) ? $info['remote_ip'] : 'no ip'),"\n";
    $server->push($request->fd, get_data(20));
});

$server->on('message', function (swoole_websocket_server $server, $frame){

});

$server->on('close', function ($server, $fd){

    unset( G::$clients[$fd] );
    echo "client {$fd} closed\n";
});

$server->start();


function ssh_connect_and_auth($name, $conf){

    $conn = ssh2_connect($conf['host'], $conf['port'], [ 'hostkey' => 'ssh-rsa' ]);
    if( !$conn ){
        echo $name, ' ', $conf['host'], " connect error\n";
        return null;
    }
    if( isset( $conf['password'] ) && $conf['password'] ){
        $ret = ssh2_auth_password($conn, $conf['username'], $conf['password']);
    }else{
        $ret = ssh2_auth_pubkey_file($conn, $conf['username'], $conf['publicKey'], $conf['privateKey']);
    }
    if( !$ret ){
        echo $name, ' ', $conf['host'], " auth error\n";
        return null;
    }

    return $conn;
}

function parse_top_data($content){

    if( !$content ){
        return null;
    }
    $rows = explode("\n", $content);
    $data = [ ];
    if( !isset( $rows[2] ) || !isset( $rows[3] ) ){
        return null;
    }
    preg_match_all('/([0-9\.]+)\%/', $rows[2], $m);
    if( !$m || !isset( $m[1][0] ) ){
        return null;
    }
    $data['cpu'] = floatval($m[1][0]);
    preg_match_all('/([0-9]+)k/', $rows[3], $m);
    if( !$m || !isset( $m[1][1] ) ){
        return null;
    }
    $data['memory'] = floatval($m[1][1]);

    return $data;
}

function get_server_info($config){

    static $conn = null;
    $data = [ ];
    foreach( $config as $name => $conf ){
        if( !isset( $conn[$name] ) || !$conn[$name] ){
            $conn[$name] = ssh_connect_and_auth($name, $conf);
            if( !$conn[$name] ){
                continue;
            }
        }
        $stream = ssh2_exec($conn[$name], 'top -n 1', 'xterm');
        if( !$stream ){
            $conn[$name] = null;
            $conn[$name] = ssh_connect_and_auth($name, $conf);
            if( !$conn['name'] ){
                continue;
            }
            $stream = ssh2_exec($conn[$name], 'top -n 1', 'xterm');
            if( !$stream ){
                echo $name, ' ', $conf['host'], " exec error\n";
                continue;
            }
        }
        stream_set_blocking($stream, true);
        $content = stream_get_contents($stream);
        fclose($stream);
        $d = parse_top_data($content);
        if( !$d ){
            echo $name, ' ', $conf['host'], " parse error\n";
            continue;
        }
        $data[$name] = $d;
    }

    return $data;
}

function put_log($type, $machine, $data, $t = null){

    $t    = $t ? : time();
    if($type === 'top'){
        $file = __DIR__ . '/data/' . $machine . '-' . date('Ymd', $t) . '.log';
        $str  = $t . ' ' . $data['cpu'] . ' ' . $data['memory'] . "\n";
    }else{
        $file = __DIR__ . '/data/' . $machine . '-' . $type . '.log';
        $str  = $t . ' ' . (is_array($data) ? implode(' ',$data) : $data) . "\n";
    }

    file_put_contents($file, $str, FILE_APPEND);
}

function get_data($count){

    $d        = [ ];
    $machine  = [ 'api', 'mysql', 'mongo', 'chat' ];
    $ts       = [ 'cpu', 'memory' ];
    if( $count != 1 ){
        $interval = 10;
        $s        = time() - ( $count - 1 ) * $interval;
        foreach( $machine as $m ){
            foreach( $ts as $v ){
                $r = null;
                for( $i = 0; $i < $count; ++$i ){
                    $r[] = [ 'x' => $s + $i * 10, 'y' => 0 ];
                }
                $d[$m][$v] = $r;
            }
        }
        for( $i = 0; $i < $count; ++$i ){
            $d['api']['request'][] = [ 'x' => $s + $i * 10, 'y' => 0 ];
        }
    }else{
        $t = time();
        $data = get_server_info(G::$sshConfig);
        foreach( $machine as $m ){
            if( isset( $data[$m] ) && $data[$m] ){
                put_log('top',$m, $data[$m], $t);
                foreach( $ts as $v ){
                    $d[$m][$v][] = [ 'x' => $t, 'y' => $data[$m][$v] ];
                }
            }else{
                foreach( $ts as $v ){
                    $d[$m][$v] = [ ];
                }
            }
        }
        $reqTime = get_api_list_timeline_request_time();
        if($reqTime === false){
            $d['api']['request'] = [];
        }else{
            put_log('request','api', $reqTime, $t);
            $d['api']['request'][] = [ 'x' => $t, 'y' => $reqTime ];
        }
    }

    return json_encode($d);
}

function request($url, $params = [], $isPost = true) {

    $opts            = [ CURLOPT_RETURNTRANSFER => true,
                         CURLOPT_HEADER         => false,
                         CURLOPT_SSL_VERIFYPEER => false,
                         CURLOPT_SSL_VERIFYHOST => false ];
    if( $isPost ){
        $opts[CURLOPT_POST]       = true;
        $opts[CURLOPT_POSTFIELDS] = http_build_query($params);
    }else{
        $url .= '?' . http_build_query($params);
    }
    $ch = curl_init($url);
    curl_setopt_array($ch, $opts);
    $ret = curl_exec($ch);
    $info = curl_getinfo($ch);
    curl_close($ch);
    if( ( false !== $ret ) && ( 200 == $info['http_code'] ) ){
        return $ret;
    }else{
        return false;
    }
}

function get_api_list_timeline_request_time(){

    if(defined('REQUEST_URL')){
        $start = microtime(true);
        if(false === request(REQUEST_URL)){
            return false;
        }
        return microtime(true) - $start;
    }

    return false;
}