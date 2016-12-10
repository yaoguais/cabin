<?php

$client = new Redis();
$client->connect('127.0.0.1', 6379);
$val = rand(10000, 99999);
$client->set('test_redis', $val);
$client->close();

sleep(1);
$client = new Redis();
$client->connect('127.0.0.1', 6380);
$ret = $client->get('test_redis');
$client->close();
echo sprintf("val:%s ret:%s\n", $val, $ret);
