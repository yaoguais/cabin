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

    curl_urls();

    sleep(5);
}


function ping_redis() {
    $host = getenv('REDIS_HOST');
    $port = getenv('REDIS_PORT');
    printf("ping host:%s port:%s\n", $host, $port);
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
    printf("ping dsn:%s user:%s password:%s\n", $dsn, $user, $password);
    $db = new PDO($dsn, $user, $password);
    $ret = $db->exec('SELECT 1');
    printf("mysql ping: %s\n", $ret);
    $db = null;
}

function curl_urls() {
    $urls = getenv('CURL_URLS');
    $urls = explode(';', $urls);
    $urls = array_filter($urls);
    foreach($urls as $url) {
        $ch = curl_init ($url);
        curl_setopt ($ch, CURLOPT_RETURNTRANSFER, 1);
        $res = curl_exec ($ch);
        curl_close ($ch);
        printf("curl %s:%s\n", $url, $res);
    }
}