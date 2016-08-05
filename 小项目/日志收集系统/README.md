# fatty

简介： fatty是一款基于Swoole的日志系统，无需修改服务端任何代码,轻松使数据存入DB中。

# 特点

1. 支持多种数据库，如MongoDb、Mysql
2. 支持多种pack方式，如php、msgpack
3. 基于Swoole TCP协议
4. 包装ThinkPHP，可以随意使用其中函数与类库
5. 使用SwooleBuffer进行数据拼接
6. 可使用nginx stream模块进行集群部署

#使用

	#启动服务器
	php server.php
	#客户端发送包
	$packageObj->addPackage($package);
	$packageObj->addPackage($package);
	$packageObj->addPackage($package);
	$packageObj->sendPackage();

# 详细步骤

1. 创建日志表
2. 修改服务端配置
3. 启动服务器
4. 编写客户端

### 1.创建日志表

	create table pre_log(
	    id int(11) not null auto_increment primary key,
	    data varchar(255) not null default ''
	);

### 2.修改服务端配置

	config.php

	return [
	    'SERVER_HOST'           => '0.0.0.0',
	    'SERVER_PORT'           => 9505,
	    'SERVER_CONFIG'         => [
	        'worker_number'  => 1,
	        'open_eof_check' => true,
	        'package_eof'    => "\r\n"
	    ],
	    'SERVER_BUFFER_CLASS' => 'fatty\\SwooleBuffer',
	    'DB_TYPE'               =>  'mysqli',
	    'DB_HOST'               =>  '127.0.0.1',
	    'DB_NAME'               =>  'test',
	    'DB_USER'               =>  'root',
	    'DB_PWD'                =>  '123456',
	    'DB_PORT'               =>  '',
	    'DB_PREFIX'             =>  'pre_',
	    'LOG_TYPE'              =>  'File'
	];

根据自身情况配置数据库，并导入上面创建的表。

### 3.启动服务器

	sudo service mysql start
	php server.php

现在就可等待客户端发送数据过来了，数据会自动保存到数据库中。

### 4.编写客户端

范例：

	require __DIR__.'/bootstrap.php';
	use \fatty\SwooleBuffer;
	use \fatty\PhpPack;
	use \fatty\MsgPack;
	use \fatty\Package;
	
	$bufferDriver = new SwooleBuffer();
	$bufferDriver->init(10240);
	$packDriver = new PhpPack();
	$packageObject = new Package();
	$packageObject->init('127.0.0.1',9505,$bufferDriver,$packDriver);
	for($i=0;$i<1000;$i++){
	    $data = array(
	        Package::DRIVER_MYSQL,
	        'log',
	        'data' => 'a log message'
	    );
	    $packageObject->addPackage($data);
	}
	echo $packageObject->sendPackage(),"\n";
	print_r($packageObject->hasError());

# Package协议

每个Package由三部分组成：驱动类型、表名/集合名、字段数据对。

	$package = array(
        Package::DRIVER_MYSQL,		//驱动类型：MYSQL
        'log',				  		//表名，不带表前缀
        'data' => 'a log message',	//字段数据1
		'data2' => 'a log message',	//字段数据2
		'data3' => 'a log message',	//字段数据3
    );

	$package = array(
        Package::DRIVER_MONGO,		//驱动类型：MONGO
        'log',				  		//集合名
        'data' => 'a log message',	//字段数据1
		'data2' => 'a log message',	//字段数据2
		'data3' => 'a log message',	//字段数据3
    );

	调用fatty\Package->addPackage($package)添加一个包
	调用fatty\Package->sendPackage()发送所有的包

# 安装

	#安装swoole
	pecl install swoole
	#安装msgpack
	pecl install msgpack
	#启动服务器
	php server.php

# ThinkPHP

fatty在对数据库进行封装的时候，并没有自己重复造轮子，而是直接将ThinkPHP这个框架改写成一个Library，框架额外占用的内存在400KB左右，启动时间增加0.02秒。

完全通过注释的方式，减少不必要的加载与路由。注释了两个文件：

1. Think.class.php
2. Mode/common.php

包装ThinkPHP的好处：一是类库丰富，二是方便项目进行扩展。

选择ThinkPHP是因为：一是本地化最好，二是超轻量级的。

# PHP VS MsgPack

	根据client.php发送的数据
	循环一千次 发送的数据长度 msg pack:31002字节 php pack:67002字节
	
	循环1000000次
	msg pack: 0.55003881454468 s
	php pack: 0.7853798866272 s
	循环1000000次
	msg pack: 0.63770604133606 s
	php pack: 0.67756915092468 s
	经过线性记录
	pack   时间: msg >= 7/5 php
	unpack 时间: msg >= 1/1 php
	长度: msg >= 2/1 php

建议使用msgpack进行数据打包。


# Fatty VS MQ

MessageQueue的麻烦之处在与Push方将数据存到Server中，Pop方并不会立即接收到通知，而只能使用poll的方式进行读取，白白浪费CPU时间。

而使用Fatty，当Push数据过来时，便可立即对数据进行相应的处理。当客户端数量与数据量很大时，一个可以增加worker的数量，还可以将数据丢到task中进行处理。