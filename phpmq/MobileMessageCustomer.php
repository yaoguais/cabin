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
	$message = $mobileMessageObj->pop();
	if(empty($message)){
		break;
	}
	$messageInfo = $message;
	send_mobile_message($messageInfo['mobile'], $messageInfo['content']);
}

echo " -- times: $num takes: ",microtime(true) - $start,"s size: ",$mobileMessageObj->size(),"\n";

function send_mobile_message($mobile,$content){
	usleep(1000);//0.001s
}