<?php

if($argc != 3){
    exit("php bench.php insertNumber selectNumber\n");
}

require(__DIR__.'/BenchBase.php');
require(__DIR__.'/mongo2.php');


$insertNumber = $argv[1] ? : 10000;
$selectNumber = $argv[2] ? : 10000;

$data = [
	'username' => 'admin',
	'password' => '123456',
	'__avpSessionId' => '0'
];

$mongoObject = new BenchMongo($data,[
    'server' => 'mongodb://192.168.154.130:27030',
    'db_name' => 'tjut',
    'collection_name' => 'user_list',
    'debug' => false
]);
$mongoObject->init();
$mongoObject->batchInset($insertNumber);
$mongoObject->repeatSelect($selectNumber);
$mongoResult = $mongoObject->getBenchMark();
unset($mongoObject);

$file = __DIR__.'/mongo_result.txt';
file_put_contents($file,$mongoResult."\r\n\r\n",FILE_APPEND);
echo file_get_contents($file);