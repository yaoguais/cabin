<?php
/**
 * RedisCluster.php
 *
 * @author    Yong Liu <newtopstdio@163.com>
 * @link      https://github.com/Yaoguais
 * @copyright 2014-2015 TJUT
 * @date      2015-07-09 17:54
 */

class RedisCluster {

    protected $db;
    protected $setSum = 0;
    protected $getSum = 0;
    protected $keys;
    protected $id;
    protected $log = '';
    protected $setArray = [];
    protected $getArray = [];

    public function __construct($host,$port){

        $this->db = new Redis();
        if(empty($this->db)){
            $this->error("extension not enable");
        }
        $this->id = 'localhost'.'_'.$port;
        $this->db->open($host,$port);//twemproxy
    }

    public function error($msg){
        exit("$msg\n");
    }


    public function batchSet($number=10000,$start=0){
        $prefix = ['tw','wt','aa','bb','cc','zz','ll','yy'];
        $max = count($prefix) - 1;
        $this->setSum = $success = 0;
        $value = str_repeat('0',1000);
        for($i=0;$i<$number;++$i){
            $key = $prefix[rand(0,$max)].($start+$i).'_'.rand(1000000,9999999);
            if($i < min($number,100)){
                $this->keys[] = $key;
            }
            $start = microtime(true);
            $ret = $this->db->set($key,$value);
            $this->setSum += (microtime(true) - $start);
            $ret===true && ++$success;
        }
        $this->setArray[] = array($number,$success,$this->setSum);
        $this->log("set:\nnumber: $number\nsuccess: $success\ntime: {$this->setSum}\n\n");
    }

    protected function log($msg){
        $this->log .= $msg;
        file_put_contents(__DIR__.'/'.$this->id.'.txt',$msg,FILE_APPEND);
    }

    public function dump(){
        echo $this->log;
        $this->log = '';
    }

    public function batchGet($number=10000){
        $this->getSum = $success = 0;
        $max = count($this->keys)-1;
        for($i=0;$i<$number;++$i){
            $key = $this->keys[rand(0,$max)];
            $start = microtime(true);
            $ret = $this->db->get($key);
            $this->getSum += (microtime(true) - $start);
            $ret!=false && ++$success;
        }
        $this->getArray[] = array($number,$success,$this->getSum);
        $this->log("get:\nnumber: $number\nsuccess: $success\ntime: {$this->getSum}\n\n");
    }

    public function export(){
        file_put_contents(__DIR__.'/'.$this->id.'.txt',var_export($this->setArray,true)."\n".var_export($this->getArray,true),FILE_APPEND);
        $avgSet = 0;
        foreach($this->setArray as $val){
            $avgSet += $val[2];
        }
        $avgSet = $avgSet/count($this->setArray);
        $avgGet = 0;
        foreach($this->getArray as $val){
            $avgGet += $val[2];
        }
        $avgGet = $avgGet/count($this->getArray);
        $number = $this->setArray[0][1] * count($this->setArray);
        file_put_contents(__DIR__.'/'.$this->id.'-avg.txt',"number: $number\nset: $avgSet\nget: $avgGet\n\n",FILE_APPEND);

        $this->setArray = $this->getArray = [];
    }

}