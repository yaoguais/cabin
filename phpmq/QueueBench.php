<?php

error_reporting(~E_WARNING & ~E_NOTICE & E_ALL);

require 'IQueue.php';
require 'RabbitQueue.php';
require 'MemcachedQueue.php';
require 'MongodbQueue.php';
require 'RedisQueue.php';
require 'MessageQueueProxy.php';

$config = require 'config.php';
$class = $config['driver'];
$mq = new $class($config['driverInfo']);
$mobileMessageObj = new MessageQueueProxy($mq);

$time = 100;
$num = 5000;

$success = 0;

for($j=0;$j<$time;++$j){
	$start = microtime(true);
	
	for($i=0;$i<$num;++$i){
		$mobile = '1355'.rand(1000000,9999999);
		$content = 'this is your phone number: '.$mobile.'.';
		$message = array(
				'mobile' => $mobile,
				'content' => $content
		);
		if(!$mobileMessageObj->push(serialize($message))){
			echo "on index:",($j*$num+$i+1)," push error !";
			goto pushout;
		}
		++$success;
	}
	$result[] = microtime(true) - $start;
	echo intval(($j+1)/$time*100),"% ";
}

pushout:

$request = $num * $time;
echo "\n";
echo "class: $class\n";
echo "push\n";
echo "times: $time num: $num\n";
echo "request: ",$request,"\n";
echo "faild: ",$request-$success,"\n";
echo "max: ",max($result)," s/{$num}times\n";
echo "min: ",min($result)," s/{$num}times\n";
echo "takes: ",$sum = array_sum($result)," s\n";
echo "average: ",$sum/count($result)," s/{$num}times\n";
echo "rqs: ",intval($time*$num/$sum),"\n";

if($success<$request){
	echo "push exit\n";
	exit();
}

$result = array();
$success = 0;
for($j=0;$j<$time;++$j){
	$start = microtime(true);

	for($i=0;$i<$num;++$i){
		if($mobileMessageObj->pop()===false){
			echo "on index:",($j*$num+$i+1)," pop error !";
			goto popout;
		}
		++$success;
	}
	$result[] = microtime(true) - $start;
	echo intval(($j+1)/$time*100),"% ";
}

popout:
echo "\n";

echo "pop\n";
echo "times: $time num: $num\n";
echo "request: ",$request,"\n";
echo "faild: ",$request-$success,"\n";
echo "max: ",max($result)," s/{$num}times\n";
echo "min: ",min($result)," s/{$num}times\n";
echo "takes: ",$sum = array_sum($result)," s\n";
echo "average: ",$sum/count($result)," s/{$num}times\n";
echo "rqs: ",intval($time*$num/$sum),"\n";
