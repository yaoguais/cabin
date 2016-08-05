<?php
/**
 * bench.php
 *
 * @author    Yong Liu <newtopstdio@163.com>
 * @link      https://github.com/Yaoguais
 * @copyright 2014-2015 TJUT
 * @date      2015-07-07 18:30
 */

require __DIR__.'/GeoBase.php';
require __DIR__.'/GeoMysql.php';

$mysqlObject = new GeoMysql([
    'dns' => 'mysql:host=127.0.0.1;port=3307;dbname=test',
    'user' => 'root',
    'password' => '123456',
]);

//$mysqlObject->createTestData(100000);

$times = [];
for($i=100;$i<2000;$i+=100){
    $t = $mysqlObject->repeatSelect($i);
    echo "$i: $t\n";
    $times[] = $t;
}

file_put_contents(__DIR__.'/times.txt',var_export($times,true),FILE_APPEND);