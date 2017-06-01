<?php
// +----------------------------------------------------------------------
// | ThinkPHP [ WE CAN DO IT JUST THINK ]
// +----------------------------------------------------------------------
// | Copyright (c) 2006-2014 http://thinkphp.cn All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: liu21st <liu21st@gmail.com>
// +----------------------------------------------------------------------

// 应用入口文件
$thinkPhpLoad['start'] = microtime(true);
// 检测PHP环境
if(version_compare(PHP_VERSION,'5.3.0','<'))  die('require PHP > 5.3.0 !');

// 开启调试模式 建议开发阶段开启 部署阶段注释或者设为false
define('APP_DEBUG',false);

//重新定义运行缓存文件目录
define('RUNTIME_PATH',__DIR__.'/');

// 定义应用目录
define('APP_PATH',__DIR__.'/Application/');

// 引入ThinkPHP入口文件
require __DIR__.'/../ThinkPHP/ThinkPHP.php';

// 亲^_^ 后面不需要任何代码了 就是如此简单

/**
 * 目前只注释掉了Think.class.php、Mode/common.php中的部分代码
 */

//设置数据库配置
$config = [
    'DB_TYPE'               =>  'mysqli',     // 数据库类型
    'DB_HOST'               =>  '127.0.0.1', // 服务器地址
    'DB_NAME'               =>  'test',          // 数据库名
    'DB_USER'               =>  'root',      // 用户名
    'DB_PWD'                =>  '123456',          // 密码
    'DB_PORT'               =>  '',        // 端口
    'DB_PREFIX'             =>  'pre_',
];
C($config);
$thinkPhpLoad['end'] = microtime(true);
echo "thinkPHP load(s):",$thinkPhpLoad['end'] - $thinkPhpLoad['start'],"\n";
/**
create table pre_log(
    id int(11) not null auto_increment primary key,
    data varchar(255) not null default ''
);
 */
/*
//数据库测试代码
$log = M('log');
for($i=0;$i<1000;++$i){
    echo $log->add(['data'=>'this is a log message']);
}
*/
/**
 * 日志系统实现细节:
 * 1. 使用msgpach序列化与反序列化(https://github.com/msgpack/msgpack-php/blob/master/msgpack.c)
 * 2. 协议细节
 *      2.1 Mysql的数据结构['表名','字段名1'=>'值1','字段名2'=>'值2']
 *      2.2 Mongo的数据结构['集合名','字段名1'=>'值1','字段名2'=>'值2']
 *      2.3 (4个字节的长度+1个字节的校验位+序列化后的数据)+"\r\n"
 *    注：由于服务器在配置的时候已经指定了数据库等（不管是集群还是什么）
 * 3.基于Swoole的TCP组件
 * 4.过滤IP地址使用操作系统级别的防火墙，而不使用应用层的IP黑名单
 */

define('CHECK_CHAR','0');

Class UdpServerLog {

    public static $cache;

    const ERROR_PACK = 1;
    const ERROR_DATA = 2;
    const ERROR_ADD = 3;
    const ERROR_LEN = 4;
    const ERROR_DIRTY_DATA = 5;
    const ERROR_CHECK = 6;

    /**
     *
        create table pre_error(
            id int(11) not null auto_increment primary key,
            dateline int(11) not null default 0,
            host varchar(23) not null default '',
            port int(11) not null default 0,
            type int(11) not null default 0,
            data varchar(255) not null default ''
        );
     * @param $host
     * @param $port
     */
    public static function addError($host,$port,$type,$data=''){
        M('error')->add([
            'dateline' => time(),
            'host' =>   $host,
            'port' => $port,
            'type' => $type,
            'data' => $data
        ]);
    }

    public static function dealMessage(&$server,$fd,&$message){
        $info = $server->connection_info($fd);
        $maxLen = strlen($message) - 2;
        echo "receive length: ".strlen($message)."\n";
        $point = 0;
        $number = $success = 0;
        while($point<$maxLen){
            ++$number;
            $packLenString = substr($message,$point,4);
            $point += 4;
            $lenArr = unpack('N',$packLenString);
            $checkChar = substr($message,$point,1);
            $point += 1;
            if($checkChar !== CHECK_CHAR){
                return UdpServerLog::addError($info['remote_ip'],$info['remote_port'],UdpServerLog::ERROR_CHECK);
            }else if(empty($lenArr) || !isset($lenArr[1])){// index begin with 1 not 0
                return UdpServerLog::addError($info['remote_ip'],$info['remote_port'],UdpServerLog::ERROR_LEN);
            }else{
                $len = $lenArr[1];
                $data = substr($message,$point,$len);
                $point += $len;
                if( false === ($data = unserialize($data)) ){
                    return UdpServerLog::addError($info['remote_ip'],$info['remote_port'],UdpServerLog::ERROR_PACK);
                }else{
                    if(!is_array($data) || count($data) < 2){
                        return UdpServerLog::addError($info['remote_ip'],$info['remote_port'],UdpServerLog::ERROR_DATA,$data);
                    }else{
                        $name = array_shift($data);
                        $model = M($name);
                        if(!$model->add($data)){
                            return UdpServerLog::addError($info['remote_ip'],$info['remote_port'],UdpServerLog::ERROR_DATA,$model->getError());
                        }
                        ++$success;
                    }
                }
            }
        }
        echo "number:{$number} success:{$success}\n";
    }
}

$server = new swoole_server('0.0.0.0', 9905, SWOOLE_PROCESS, SWOOLE_SOCK_TCP);
$server->set([
    'open_eof_check' => true, //打开EOF检测
    'package_eof' => "\r\n"
]);
$server->on('receive', function (swoole_server $server, $fd, $from_id, $message)
{
    echo "receive length: ".strlen($message)."\n";
    $oef = substr($message,-2);
    if($oef!=="\r\n"){
        if(isset(UdpServerLog::$cache[$fd])){
            UdpServerLog::$cache[$fd] .= $message;
        }else{
            UdpServerLog::$cache[$fd] = $message;
        }
    }else{
        if(isset(UdpServerLog::$cache[$fd])){
            $message = UdpServerLog::$cache[$fd].$message;
            UdpServerLog::$cache[$fd] = null;
            UdpServerLog::dealMessage($server,$fd,$message);
        }else{
            UdpServerLog::dealMessage($server,$fd,$message);
        }
    }
});

$server->on('close',function(swoole_server $server, $fd,$from_id){
    if(isset(UdpServerLog::$cache[$fd]) && UdpServerLog::$cache[$fd]){
        UdpServerLog::$cache[$fd] = null;
        $info = $server->connection_info($fd);
        UdpServerLog::addError($info['remote_ip'],$info['remote_port'],UdpServerLog::ERROR_DIRTY_DATA,$fd);
    }
});
$server->start();