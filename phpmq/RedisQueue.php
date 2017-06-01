<?php

class RedisQueue implements IQueue{
	
	private $_conn = null;
	private $_connInfo = null;
	
	public function __construct($connInfo){
		if(empty($connInfo['key'])){
			$connInfo['key'] = 'rmq';
		}
		$this->_connInfo = $connInfo;
	}
	
	public function open(){
		if(empty($this->_connInfo)){
			return false;
		}
		if($this->_connInfo['newlink']){
			$this->close();
		}
		if($this->_conn){
			return true;
		}
		$this->_conn = new Redis();
		if(empty($this->_conn)){
			return false;
		}
		if($this->_connInfo['persistent']){
			if(!$this->_conn->pconnect($this->_connInfo['host'],$this->_connInfo['port'])){
				return false;
			}
		}else{
			if(!$this->_conn->connect($this->_connInfo['host'],$this->_connInfo['port'])){
				return false;
			}
		}
		return true;
	}
	
	public function push($message){
		if(!$this->open()){
			return false;
		}
		return $this->_conn->rPush($this->_connInfo['key'],$message);
	}
	
	public function pop(){
		if(!$this->open()){
			return false;
		}
		return $this->_conn->lPop($this->_connInfo['key']);
	}
	
	public function front(){
		if(!$this->open()){
			return false;
		}
		return $this->_conn->lGet($this->_connInfo['key'],0);
	}
	
	public function back(){
		if(!$this->open()){
			return false;
		}
		return $this->_conn->lGet($this->_connInfo['key'],-1);
	}
	
	public function isEmpty(){
		if(!$this->open()){
			return true;
		}
		return $this->_conn->lLen($this->_connInfo['key'])>0;
	}
	
	public function size(){
		if(!$this->open()){
			return -1;
		}
		return $this->_conn->lLen($this->_connInfo['key']);
	}
	
	public function close(){
		if($this->_conn){
			$this->_conn->close();
			$this->_conn = null;
		}
	}
}