<?php

$pdo = new PDO('mysql:host=127.0.0.1;port=3306', 'root', 'yaoguais_2014');
$sql = <<<EOF
-- create database if not exists test;
use test;
create table if not exists test_user(
    uid int,
    username varchar(32)
);
insert into test_user(username) values('yaoguais');
EOF;
$sqlList = explode(";\n", $sql);
foreach ($sqlList as $sql) {
    if (trim($sql)) {
        $pdo->exec($sql);
    }
}
$stmt = $pdo->query('select count(*) as cnt from test_user');
$ret = $stmt->fetchAll(PDO::FETCH_ASSOC);
echo "master:\n";
print_r($ret);
$pdo = null;

$pdo = new PDO('mysql:dbname=test;host=127.0.0.1;port=3307', 'root', 'yaoguais_2014');
$stmt = $pdo->query('select count(*) as cnt from test_user');
$ret = $stmt->fetchAll(PDO::FETCH_ASSOC);
echo "slave:\n";
print_r($ret);
$pdo = null;
