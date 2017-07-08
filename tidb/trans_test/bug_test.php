<?php

// Copy from laravel
$options = $options = [
    PDO::ATTR_CASE => PDO::CASE_NATURAL,
    PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
    PDO::ATTR_ORACLE_NULLS => PDO::NULL_NATURAL,
    PDO::ATTR_STRINGIFY_FETCHES => false,
    PDO::ATTR_EMULATE_PREPARES => false,
];

$pdo = new PDO('mysql:host=192.168.12.254;port=4000;dbname=test', 'root', '', $options);
// $pdo = new PDO('mysql:host=127.0.0.1;dbname=test', 'root', '', $options);

// Test table
// create table test (id bigint(20) auto_increment primary key, stock int(10) unsigned not null default 0);

// First
// insert into test(id, stock) values(1, 10);
// or update test set stock = 10 where id = 1;


// Second, decrement stock and insert a value
$ret = $pdo->beginTransaction();
file_log(sprintf('begin ret %s', $ret));
if ($ret) {
    try {
        $ret = $pdo->exec('update test set stock = stock - 10 where id = 1');
        if (!$ret) {
            throw new \Exception('update stock failed');
        }
        $ret = $pdo->exec(sprintf('insert into test(stock) values(%d)', rand(0, 9999)));
        if (!$ret) {
            throw new Exception('insert failed');
        }
        $id = $pdo->lastInsertId();
        file_log(sprintf('last insertId %d', $id));
        $ret = $pdo->commit();
        file_log(sprintf('commit ret %d', $ret));
        $res = $pdo->query(sprintf('select * from test where id = %d', $id));
        $row = $res->fetch(PDO::FETCH_ASSOC);
        file_log(sprintf('inserted data %s', json_encode($row)));
    } catch (Exception $e) {
        $ret = $pdo->rollBack();
        file_log(sprintf('rollback %d exception %s', $ret, $e->getTraceAsString()));
    } catch (Throwable $e) {
        $ret = $pdo->rollBack();
        file_log(sprintf('rollback %d exception %s', $ret, $e->getTraceAsString()));
    }
}

// Third, one terminal 'php -S 127.0.0.1:9093'
//        another terminal 'ab -c 10 -n 10 http://127.0.0.1:9093/bug_test.php'

function file_log($msg)
{
    static $rId = null;
    if (is_null($rId)) {
        $rId = uniqid();
    }
    file_put_contents(__DIR__ . '/test.log', sprintf("%s %s\n", $rId, $msg), FILE_APPEND);
}
