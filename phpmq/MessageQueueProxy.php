<?php

class MessageQueueProxy{

	private $_mq;

	public function __construct($mq){
		$this->_mq = $mq;
	}

	public function push($mesage){
		return $this->_mq->push($mesage);
	}

	public function pop(){
		return $this->_mq->pop();
	}
	
	/*仅用作测试*/
	public function size(){
		return $this->_mq->size();
	}

}