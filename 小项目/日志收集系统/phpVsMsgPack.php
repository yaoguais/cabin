<?php
/**
 * Created by PhpStorm.
 * User: yaoguai
 * Date: 15-7-5
 * Time: 下午10:37
 */
require __DIR__.'/bootstrap.php';
use \fatty\PhpPack;
use \fatty\MsgPack;
use \fatty\Package;

$data = array(
    Package::DRIVER_MYSQL,
    'log',
    'data' => 'a log message'
);
$phpPack = new PhpPack();
$msgPack = new MsgPack();
$number = 1000000;
echo "number: {$number}\n";
$start = microtime(true);
for($i=0;$i<$number;++$i){
    $phpPack->pack($data);
}
echo "php pack ",(microtime(true) - $start)," s\n";
$start = microtime(true);
for($i=0;$i<$number;++$i){
    $msgPack->pack($data);
}
echo "msg pack ",(microtime(true) - $start)," s\n";
echo "php pack len: ",(strlen($phpString = $phpPack->pack($data)))," byte\n";
echo "msg pack len: ",(strlen($msgString = $msgPack->pack($data)))," byte\n";

$number = 1000000;
echo "number: {$number}\n";
$start = microtime(true);
for($i=0;$i<$number;++$i){
    $phpPack->unpack($phpString);
}
echo "php unpack ",(microtime(true) - $start)," s\n";
$start = microtime(true);
for($i=0;$i<$number;++$i){
    $msgPack->unpack($msgString);
}
echo "msg unpack ",(microtime(true) - $start)," s\n";



/**
 * 根据client.php发送的数据
 * 循环一千次 发送的数据长度 msg pack:31002字节 php pack:67002字节
 *
 * 循环1000000次
 * msg pack: 0.55003881454468 s
 * php pack: 0.7853798866272 s
 * 循环1000000次
 * msg pack: 0.63770604133606 s
 * php pack: 0.67756915092468 s
 * 经过线性记录
 * pack   时间: msg >= 7/5 php
 * unpack 时间: msg >= 1/1 php
 * 长度: msg >= 2/1 php
**/



