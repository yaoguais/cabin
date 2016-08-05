<?php
/**
 * bench.php
 *
 * @author    Yong Liu <newtopstdio@163.com>
 * @link      https://github.com/Yaoguais
 * @copyright 2014-2015 TJUT
 * @date      2015-07-07 10:47
 */

if($argc != 3){
    exit("php bench.php insertNumber selectNumber\n");
}

require(__DIR__.'/BenchBase.php');
require(__DIR__.'/mysql.php');
require(__DIR__.'/mongo.php');


$insertNumber = $argv[1] ? : 10000;
$selectNumber = $argv[2] ? : 10000;
/**
CREATE TABLE bench_test(
    id int(11) not null auto_increment primary key,
    name varchar(255) not null default '',
    friend varchar(255) not null default '',
    relationship varchar(255) not null default '',
    note varchar(255) not null default ''
)engine = INNODB DEFAULT CHARSET UTF8;
 */

$data = [
    'name' => 'liMeiMei',
    'friend' => 'hanLeiLei',
    'relationship' => 'english friend',
    'note' => 'this is a funny girl'
];

$mysqlObject = new BenchMysql($data,[
    'dns' => 'mysql:host=localhost;dbname=test',
    'user' => 'root',
    'password' => '123456',
    'table' => 'bench_test',
    'debug' => false
]);
$mysqlObject->init();
$mysqlObject->batchInset($insertNumber);
$mysqlObject->repeatSelect($selectNumber);
$mysqlResult = $mysqlObject->getBenchMark();
unset($mysqlObject);
$mongoObject = new BenchMongo($data,[
    'server' => 'mongodb://localhost:27017',
    'db_name' => 'benchTest',
    'collection_name' => 'bench_test',
    'debug' => false
]);
$mongoObject->init();
$mongoObject->batchInset($insertNumber);
$mongoObject->repeatSelect($selectNumber);
$mongoResult = $mongoObject->getBenchMark();
unset($mongoObject);

$file = __DIR__.'/result.txt';
file_put_contents($file,$mysqlResult."\r\n".$mongoResult."\r\n\r\n",FILE_APPEND);
echo file_get_contents($file);


