<?php

//参考http://www.jb51.net/article/50286.htm实现
//该实现是用两个指针,分别指向队首与队尾,然后对队列进行操作.
//这不具备原子性,在多进程下会导致数据混乱的情况
//但是我们还是实现一下,看看其性能怎么样

class MemcachedQueue implements IQueue{
	private $_conn = null;
	private $_connInfo = null;
	private $_key = 'mmq';
	
	public function __construct($connInfo){
		if($connInfo['key']){
			$this->_key = $connInfo['key'];
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
		if($this->_connInfo['persistent_id ']){
			$this->_conn = new Memcached($this->_connInfo['persistent_id ']);
		}else{
			$this->_conn = new Memcached();
		}
		if(empty($this->_conn)){
			return false;
		}
		if($this->_connInfo['servers']){
			$this->_conn->addServers($this->_connInfo['servers']);
		}else{
			$this->_conn->addServer($this->_connInfo['host'],$this->_connInfo['port'],$this->_connInfo['weight'] ? : 0);
		}
		return true;
	}
	
	private function _setCounter( $key, $offset, $time=0 ){
		if(!$this->open()){
			return false;
		}
		$val = $this->_conn->get($key);
		if( !is_numeric($val) || $val < 0 ){
			$ret = $this->_conn->set( $key, 0, $time );
			if( !$ret ) return false;
			$val = 0;
		}
		$offset = intval( $offset );
		if( $offset > 0 ){
			return $this->_conn->increment( $key, $offset );
		}elseif( $offset < 0 ){
			return $this->_conn->decrement( $key, -$offset );
		}
		return $val;
	}
	
	public function push($message){
		if(!$this->open()){
			return false;
		}
		$pushKey = $this->_key.'w';
		if(false === ($pushIndex = $this->_setCounter($pushKey, 1))){
			return false;
		}
		$valueKey = $this->_key.$pushIndex;
		return $this->_conn->set($valueKey,$message);
	}
	
	public function pop(){
		if(!$this->open()){
			return false;
		}	
		$pushKey = $this->_key.'w';
		if(false === ($pushIndex = $this->_setCounter($pushKey, 0))){
			return false;
		}
		$popKey = $this->_key.'r';
		if(false === ($popIndex = $this->_setCounter($popKey, 0))){
			return false;
		}
		++$popIndex;
		if($pushIndex < $popIndex){
			return false;
		}
		$valueKey = $this->_key.$popIndex;
		if(false === ($this->_setCounter($popKey, 1))){
			return false;
		}
		$ret = $this->_conn->get($valueKey);
		if(empty($ret)){
			return false;
		}
		$this->_conn->delete($valueKey);
		return $ret;
	}
	
	public function front(){
		if(!$this->open()){
			return false;
		}
		$pushKey = $this->_key.'w';
		$pushIndex = $this->_setCounter($pushKey, 0);
		$valueKey = $this->_key.$pushIndex;
		return $this->_conn->get($valueKey);
	}
	
	public function back(){
		if(!$this->open()){
			return false;
		}
		$popKey = $this->_key.'r';
		$popIndex = $this->_setCounter($popKey, 0);
		$valueKey = $this->_key.$popIndex;
		return $this->_conn->get($valueKey);
	}
	
	public function isEmpty(){
		if(!$this->open()){
			return true;
		}
		return $this->size()>0;
	}
	
	public function size(){
		if(!$this->open()){
			return -1;
		}
		$pushKey = $this->_key.'w';
		$pushIndex = $this->_setCounter($pushKey, 0);
		$popKey = $this->_key.'r';
		$popIndex = $this->_setCounter($popKey, 0);
		return $pushIndex - $popIndex;
	}
	
	public function close(){
		if($this->_conn){
			$this->_conn->quit();
			$this->_conn = null;
		}
	}
	
	private function _debug($msg){
		$pushKey = $this->_key.'w';
		$pushIndex = $this->_setCounter($pushKey, 0);
		$popKey = $this->_key.'r';
		$popIndex = $this->_setCounter($popKey, 0);
		print_r($this->_conn->getStats());
		echo "$msg pushIndex: $pushIndex / popIndex: $popIndex\n";
	}
}