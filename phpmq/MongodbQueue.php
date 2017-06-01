<?php

/**
 * 使用timestamp实现queue
 * mongodb还有一种使用findAndModify实现队列,但是出队并不删除元素,这个随着时间的推移,内存必爆。
 * 所以这里就不实现这种了。
 * timestamp各种关系性数据库都能实现,这里也不实现mysql等数据库的了.
 * @author yaoguai
 *
 */

class MongodbQueue implements IQueue{
	
	private $_conn;
	private $_connInfo;
	
	/**
	 * 返回当前的毫秒数
	 * @return number
	 */
	private function _getTimeStamp(){
		return intval(microtime(true)*1000);
	}
	
	public function __construct($connInfo){
		if(empty($connInfo['db_name'])){
			$connInfo['db_name'] = 'mmq';
		}
		if(empty($connInfo['collection_name'])){
			$connInfo['collection_name'] = 'cmmq';
		}
		$this->_connInfo = $connInfo;
	}
	
	public function open(){
		if($this->_conn){
			return true;
		}
		if(!class_exists('MongoClient')){
			return false;
		}
		$this->_conn = new MongoClient($this->_connInfo['server'],$this->_connInfo['server_options']);
		if(empty($this->_conn)){
			return false;
		}
		//初始化数据库，集合，索引等
		$collection = $this->_conn->{$this->_connInfo['db_name']}->{$this->_connInfo['collection_name']};
		$collection->createIndex(array('timestamp'=>1));
		return true;
	}
	
	public function close(){
		if($this->_conn){
			$this->_conn->close();
		}
	}
	
	public function push($message){
		if(!$this->open()){
			return false;
		}
		$arr = array(
				'timestamp' => $this->_getTimeStamp(),
				'message' => $message
		);
		$collection = $this->_conn->{$this->_connInfo['db_name']}->{$this->_connInfo['collection_name']};		
		$ret = $collection->insert($arr);
		return isset($ret['ok']) ? ($ret['ok'] ? true : false) : ($ret ? true : false);
	}
	
	public function pop(){
		if(!$this->open()){
			return false;
		}
		//首先查询这个集合然后删除
		$collection = $this->_conn->{$this->_connInfo['db_name']}->{$this->_connInfo['collection_name']};
		$cursor = $collection->find()->sort(array('timestamp'=>1))->limit(1);
		$ret = $cursor->getNext();
		if(empty($ret)){
			return false;
		}
		/*这里只比较ID*/
		if($collection->remove(array('_id'=>$ret['_id']),array("justOne"=>true))){
			return $ret;
		}else{
			return false;
		}
	}
	
	public function front(){
		if(!$this->open()){
			return false;
		}
		$collection = $this->_conn->{$this->_connInfo['db_name']}->{$this->_connInfo['collection_name']};
		$cursor = $collection->find()->sort(array('timestamp'=>1))->limit(1);
		$ret = $cursor->getNext();
		return $ret ? : false;
	}
	
	public function back(){
		if(!$this->open()){
			return false;
		}
		$collection = $this->_conn->{$this->_connInfo['db_name']}->{$this->_connInfo['collection_name']};
		$cursor = $collection->find()->sort(array('timestamp'=>-1))->limit(1);
		$ret = $cursor->getNext();
		return $ret ? : false;
	}
	
	public function isEmpty(){
		if(!$this->open()){
			return true;
		}
		$collection = $this->_conn->{$this->_connInfo['db_name']}->{$this->_connInfo['collection_name']};
		$cursor = $collection->find();
		return $cursor->count()>0;
	}
	
	public function size(){
		if(!$this->open()){
			return -1;
		}
		$collection = $this->_conn->{$this->_connInfo['db_name']}->{$this->_connInfo['collection_name']};
		$cursor = $collection->find();
		return $cursor->count();
	}
}