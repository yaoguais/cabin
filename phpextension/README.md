# monster
a simple php class by c extension

## 类定义 ##

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


## 测试代码 ##

	//construct
	$obj = new monster(array('people','animal'));
	//getFoods
	var_dump($obj->getFoods());
	//eat
	var_dump($obj->eat('stone'));
	//eat
	var_dump($obj->eat('animal'));
	//addFood
	$obj->addFood('meat');
	$obj->addFood('animal');
	$obj->addFood('animal');
	//getFoods
	var_dump($obj->getFoods());
	//delFood
	$obj->delFood('animal');
	//getFoods
	var_dump($obj->getFoods());


## 结语 ##

就是实现一个简单的类，并没有实际意义的类。

但是这也为实现复杂的功能打下了基础。
