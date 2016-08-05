<?php 
class monster {
		
	private $_foods = array();
	
	public function __construct($foods){
		$this->_foods = $foods;
	}
	
	public function eat($food){
		if(in_array($food,$this->_foods)){
			return true;
		}else{
			return false;
		}
	}

	public function addFood($food){
		$this->_foods[] = $food;
	}
	
	public function delFood($food){
		$this->_foods[$food] = null;
	}
	
	public function getFoods(){
		return $this->_foods();
	}
	
}