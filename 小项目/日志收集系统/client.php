<?php
/**
 * Created by PhpStorm.
 * User: yaoguai
 * Date: 15-7-5
 * Time: 下午5:35
 */

require __DIR__.'/bootstrap.php';
use \fatty\SwooleBuffer;
use \fatty\PhpPack;
use \fatty\MsgPack;
use \fatty\Package;

$bufferDriver = new SwooleBuffer();
$bufferDriver->init(10240);
$packDriver = new PhpPack();
$packageObject = new Package();
$packageObject->init('127.0.0.1',9505,$bufferDriver,$packDriver);
for($i=0;$i<10;$i++){
    $data = array(
        Package::DRIVER_MYSQL,
        'log',
        'data' => 'a log message'
    );
    $packageObject->addPackage($data);
}
echo $packageObject->sendPackage(),"\n";
print_r($packageObject->hasError());