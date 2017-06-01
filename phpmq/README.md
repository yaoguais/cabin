# PHP消息队列的一些实现

目录：

1. 消息队列的简介
2. 性能测试


## 消息队列的简介 ##

消息队列（Message Queue）：把消息按照产生的次序加入队列，而由另外的处理程序/模块将其从队列中取出，并加以处理；从而形成了一个基本的消息队列。应用场景：短信服务、电子邮件服务、图片处理服务、好友动态推送服务等。
[更多可以看看这篇文章](http://yaoguais.github.io/?s=md/php/mq.md)

这里主要做了以下几个实现：

1. redis+list
2. mongodb+index+timestamp
3. memcached+pointer
4. amqp+rabbitmq

## 性能测试 ##

性能测试使用的是QueueBench.php实现的，其环境参数与测试结果如下:

	php: 5.4.40
	Linux: 3.13.0-34-generic(buildd@allspice) (gcc version 4.8.2 (Ubuntu 4.8.2-19ubuntu1) )
	memory: 4GB
	cpu: Intel(R)_Core(TM)_i5@3.00GHz*4
	redis: 3.0.1
	php-redis: 2.2.7
	mongodb: mongodb-linux-x86_64-ubuntu1404-3.0.2
	php-mongo: 1.6.7
	memcached: 1.4.24
	php-memcached: 2.2.0

redis的数据如下：

	class: RedisQueue
	push
	times: 100 num: 10000
	request: 1000000
	faild: 0
	max: 0.37430000305176 s/10000times
	min: 0.26855707168579 s/10000times
	takes: 32.56348824501 s
	average: 0.3256348824501 s/10000times
	rqs: 30709
	pop
	times: 100 num: 10000
	request: 1000000
	faild: 0
	max: 0.28984785079956 s/10000times
	min: 0.21943783760071 s/10000times
	takes: 26.595669269562 s
	average: 0.26595669269562 s/10000times
	rqs: 37600

mongodb+index的数据如下：

	class: MongodbQueue
	push
	times: 100 num: 10000
	request: 1000000
	faild: 0
	max: 1.8860659599304 s/10000times
	min: 1.038162946701 s/10000times
	takes:115.00487303734 s
	average: 1.1500487303734 s/10000times
	rqs: 8695
	pop
	times: 100 num: 10000
	request: 1000000
	faild: 0
	max: 2.2940940856934 s/10000times
	min: 1.8466329574585 s/10000times
	takes: 197.70558190346 s
	average: 1.9770558190346 s/10000times
	rqs: 5058

mongodb+noindex的数据如下:

	class: MongodbQueue
	push
	times: 100 num: 10000
	request: 1000000
	faild: 0
	max: 1.4391729831696 s/10000times
	min: 1.0708310604095 s/10000times
	takes: 120.37140989304 s
	average: 1.2037140989304 s/10000times
	rqs: 8307
	//下面根本跑不动了

memcached的数据如下:

	class: MemcachedQueue
	push
	times: 100 num: 10000
	request: 1000000
	faild: 0
	max: 0.83561706542969 s/10000times
	min: 0.67476201057434 s/10000times
	takes: 76.049747467041 s
	average: 0.76049747467041 s/10000times
	rqs: 13149
	pop
	times: 100 num: 10000
	request: 1000000
	faild: 0
	max: 1.3662950992584 s/10000times
	min: 1.2345328330994 s/10000times
	takes: 129.71623921394 s
	average: 1.2971623921394 s/10000times
	rqs: 7709

rabbitmq的数据如下:

	class: RabbitQueue
	push
	times: 100 num: 5000
	request: 500000
	faild: 0
	max: 0.56668591499329 s/5000times
	min: 0.019957065582275 s/5000times
	takes: 17.697237014771 s
	average: 0.17697237014771 s/5000times
	rqs: 28252
	pop
	times: 100 num: 5000
	request: 500000
	faild: 0
	max: 1.709969997406 s/5000times
	min: 0.56408095359802 s/5000times
	takes: 85.071734666824 s
	average: 0.85071734666824 s/5000times
	rqs: 5877

	VMware 10
	centos 6.6
	memory 2GB
	cpu i5*4

关于rabbitmq的更多信息，[请看这里](http://yaoguais.github.io/?s=md/php/mq.md#消息队列的实现amqp-rabbitmq)

可以看出：

- 可以看出redis在push上是mongodb的3.53倍,在pop上是mongodb的7.43倍.
- mongodb没有索引,pop根本跑不动了。过了10多秒，我进数据库一看，才删除了99条.
- 可以看出redis的push是memcached的2.3倍,redis的pop是memcached的4.5倍.
- rabbitmq性能并不是最优的
