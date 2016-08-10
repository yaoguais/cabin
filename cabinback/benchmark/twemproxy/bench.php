<?php
/**
 * bench.php
 *
 * @author    Yong Liu <newtopstdio@163.com>
 * @link      https://github.com/Yaoguais
 * @copyright 2014-2015 TJUT
 * @date      2015-07-09 18:33
 */

require __DIR__.'/RedisCluster.php';

$number = 100000;
$loop = 10;
$twemproxy = new RedisCluster('127.0.0.1',22102);
$objArr = [$twemproxy];

foreach($objArr as $obj){
    for($i=0;$i<$loop;++$i){
        $start = $i*$number;
        $obj->batchSet($number,$start);
        $obj->batchGet($number);
        $obj->dump();
    }
    $obj->export();
}