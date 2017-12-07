<?php

while(true) {
    try {
        ping_redis();
    } catch(Exception $e) {
        printf("exception: %s\n", $e->getMessage());
    }

    try {
        ping_mysql();
    } catch(Exception $e) {
        printf("exception: %s\n", $e->getMessage());
    }

    sleep(5);
}


function ping_redis() {
    $host = getenv('REDIS_HOST');
    $port = getenv('REDIS_PORT');
    $redis = new Redis();
    $redis->connect($host, $port);
    $ret = $redis->ping();
    printf("redis ping: %s\n", $ret);
    $redis->close();
}

function ping_mysql() {
    $host = getenv('MYSQL_HOST');
    $port = getenv('MYSQL_PORT');
    $user = getenv('MYSQL_USER');
    $password  = getenv('MYSQL_PASSWORD');
    $database = getenv('MYSQL_DATABASE');
    $charset = 'utf8';

    $dsn = sprintf('mysql:host=%s;port=%s;dbname=%s;charset=%s', $host, $port, $database, $charset);
    $db = new PDO($dsn, $user, $password);
    $ret = $db->exec('SELECT 1');
    printf("mysql ping: %s\n", $ret);
    $db = null;
}
