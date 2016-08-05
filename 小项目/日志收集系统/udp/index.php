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
 *    注：由于服务器在配置的时候已经指定了数据库等（不管是集群还是什么）
 * 3.基于Swoole的UDP组件
 * 4.过滤IP地址使用操作系统级别的防火墙，而不使用应用层的IP黑名单
 */

Class UdpServerLog {

    public static $count = 0;
    public static $lastUpdate = 0;

    const ERROR_PACK = 1;
    const ERROR_DATA = 2;
    const ERROR_ADD = 3;

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
}

$server = new swoole_server('0.0.0.0', 9905, SWOOLE_PROCESS, SWOOLE_SOCK_UDP);
$server->on('packet', function (swoole_server $serv, $data, $addr)
{
    if(time() - UdpServerLog::$lastUpdate > 10){
        UdpServerLog::$count = 0;
        echo "reset:----------------\n";
    }
    UdpServerLog::$lastUpdate = time();
    ++UdpServerLog::$count;
    file_put_contents('receive-count.txt',UdpServerLog::$count);
    //$serv->sendto($addr['address'], $addr['port'], "Swoole: $data");
    if( null === ($data = unserialize($data)) ){
        UdpServerLog::addError($addr['address'], $addr['port'],UdpServerLog::ERROR_PACK);
    }else{
        if(!is_array($data) || count($data) < 2){
            UdpServerLog::addError($addr['address'], $addr['port'],UdpServerLog::ERROR_DATA,$data);
        }else{
            $name = array_shift($data);
            $model = M($name);
            if(!$model->add($data)){
                UdpServerLog::addError($addr['address'], $addr['port'],UdpServerLog::ERROR_DATA,$model->getError());
            }
        }
    }
});
$server->start();