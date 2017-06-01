<?php

class RabbitQueue implements IQueue {

    private $_conn;
    private $_connInfo;
    private $_exchange;
    private $_queue;

    function __construct($connInfo) {
        if(empty($connInfo['routingKey'])){
            $connInfo['routingKey'] = 'routingKey';
        }
        $connInfo['bindingkey'] = $connInfo['routingKey'];
        if(empty($connInfo['exchangeName'])){
            $connInfo['exchangeName'] = 'exchangeName';
        }
        if(empty($connInfo['queueName'])){
            $connInfo['queueName'] = 'queueName';
        }
        $this->_connInfo = $connInfo;
    }

    public function open(){
        if(empty($this->_connInfo)){
            return false;
        }
        if($this->_conn){
            return true;
        }
        $this->_conn = new AMQPConnection($this->_connInfo['connArgs']);
        if (!$this->_conn->connect()) {
            return false;
        }
        return true;
    }

    private function _initPublisher(){
        if(!$this->open()){
            return false;
        }
        if($this->_exchange){
            return true;
        }
        $channel = new AMQPChannel($this->_conn);
        //创建exchange
        $this->_exchange = new AMQPExchange($channel);
        $this->_exchange->setName($this->_connInfo['exchangeName']);
        $this->_exchange->setType(AMQP_EX_TYPE_DIRECT);
        $this->_exchange->setFlags(AMQP_DURABLE);
        if(!$this->_exchange->declareExchange()) {
            $this->close();
            return false;
        }
        //创建队列,生产者和消费者都要创建队列
        $queue = new AMQPQueue($channel);
        $queue->setName($this->_connInfo['queueName']);
        $queue->setFlags(AMQP_DURABLE);
        $queue->declareQueue();
        $queue->bind($this->_connInfo['exchangeName'],$this->_connInfo['routingKey']);
        return true;
    }

    private function _initCustomer(){
        if(!$this->open()){
            return false;
        }
        if($this->_queue){
            return true;
        }
        $channel = new AMQPChannel($this->_conn);
        $this->_queue = new AMQPQueue($channel);
        if(!$this->_queue){
            $this->close();
            return false;
        }
        $this->_queue->setName($this->_connInfo['queueName']);
        $this->_queue->setFlags(AMQP_DURABLE);
        $this->_queue->declareQueue();
        $this->_queue->bind($this->_connInfo['exchangeName'],$this->_connInfo['bindingKey']);
        return true;
    }

    public function close(){
        if($this->_conn){
            $this->_conn->disconnect();
            $this->_conn = null;
        }
    }

    public function push($message) {
        if(!$this->open() || !$this->_initPublisher()){
            return false;
        }
        return $this->_exchange->publish($message,$this->_connInfo['routingKey']);
    }


    public function pop() {
        if(!$this->open() || !$this->_initCustomer()){
            return false;
        }
        $messages = $this->_queue->get(AMQP_AUTOACK);
        if(empty($messages)){
            return false;
        }
        return $messages->getBody();
    }

    public function front() {
        return false;
    }

    public function back() {
        return false;
    }

    public function isEmpty() {
        return true;
    }

    public function size() {
        return -1;
    }

}
