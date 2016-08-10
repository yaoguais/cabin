<?php
/*
CREATE TABLE IF NOT EXISTS user(
  uid int(11) not null auto_increment primary key,
  username varchar(32) not null default '',
  description varchar(1000) not null default ''
);
*/

$mysqlObject = new PDO('mysql:host=localhost;dbname=tjut','root','123456');
$mongoObject = new MongoClient('mongodb://192.168.1.131:27000');
$redisObject = new Redis();

//查询最新的用户ID
function get_max_uid($db){
    $query = $db->query('SELECT uid from user order by uid desc limit 1',PDO::FETCH_ASSOC);
    $result = $query->fetch();
    return $result ? $result['uid'] : 1;
}

function get_user($db,$uid){
    $query = $db->query('SELECT * from user where uid = '.$uid,PDO::FETCH_ASSOC);
    return $query->fetch($query);
}

if(isset($_GET['insert'])){
    $minUid = get_max_uid($mysqlObject);
    $description = str_repeat('A',900);
    $mongoObject->selectDB('tjut');
    for($i=1;$i<=1000;++$i){
        $uid = $i + $minUid;
        $username = 'admin'.$uid;
        $desc = $username.$description;
        //MySQL插入
        $mysqlObject->exec('INSERT INTO user(username,description) values(\''.$username.'\',\''.$desc.'\');');
        //Mongo写插入成功
        $mongoObject->selectCollection('tjut','user')->insert([
            'uid' => $uid,
            'username' => $username,
            'description' => $desc
        ]);
    }
    echo get_max_uid($mysqlObject),"\n";
}else{
    $maxUid = get_max_uid($mysqlObject);
    $cacheTime = 0;
    $dbTime = 0;
    $redisObject->open('127.0.0.1',22100);
    for($i=0;$i<1000;++$i){
        $uid = rand(1,$maxUid);
        $cacheValue = $redisObject->get($uid);
        if(empty($cacheValue)){
            $user = get_user($mysqlObject,$uid);
            $redisObject->set($uid,serialize($user));
            ++$dbTime;
        }else{
            $user = unserialize($cacheValue);
            ++$cacheTime;
        }
    }
    echo "db:{$dbTime} cache:{$cacheTime}\n";
}

