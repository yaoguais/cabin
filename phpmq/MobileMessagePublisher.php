<?php

error_reporting(~E_NOTICE & E_ALL);

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

$num = 10000;
echo $start = microtime(true);

for($i=0;$i<$num;++$i){
	$mobile = '1355'.rand(1000000,9999999);
	$content = 'this is your phone number: '.$mobile.'.';
	$message = array(
			'mobile' => $mobile,
			'content' => $content
	);
	if(!$mobileMessageObj->push(serialize($message))){
		echo "push error !";
		break;
	}
}

echo " -- times: $num takes: ",microtime(true) - $start,"s size: ",$mobileMessageObj->size(),"\n";