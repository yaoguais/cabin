<?php
/**
 * mongo.php
 *
 * @author    Yong Liu <newtopstdio@163.com>
 * @link      https://github.com/Yaoguais
 * @copyright 2014-2015 TJUT
 * @date      2015-07-07 10:47
 */

class BenchMongo extends BenchBase{

    protected $collection;

    public function __construct($data,$config){
        $this->data = $data;
        $this->config = $config;
        $this->driver = 'mongo';
    }

    public function init(){
        $this->insert = 0;
        $this->select = 0;
        $this->error = 0;
        if(!class_exists('MongoClient')){
            $this->error('mongo is not enabled');
        }
        $c = & $this->config;
        $this->db = new MongoClient($c['server']);
        if(empty($this->db)){
            $this->error('mongo connect error');
        }
        $this->collection = $this->db->{$c['db_name']}->{$c['collection_name']};
    }

    public function insert(){
        $data = $this->data;
        $c = & $this->config;
        foreach($data as $key=>&$val){
            $val .= rand(1000000,9999999);
            $c['key'] = $key;
            $c['val'] = $val;
        }
        $start = microtime(true);
        $ret = $this->collection->insert($data);
        $this->insert += (microtime(true) - $start);
        if(empty($ret)){
            ++$this->error;
            if($c['debug']){
                echo "insert error\n";
            }
        }
        if($c['debug']){
            print_r($this->db->{$c['db_name']}->lastError());
        }
    }

    public function select(){
        $c = & $this->config;
        $query = [
            $c['key'] => $c['val']
        ];
        $start = microtime(true);
        $cursor =  $this->collection->find($query);
        if(empty($cursor)){
            ++$this->error;
            if($c['debug']){
                echo "find error\n";
            }
            return;
        }
        $result = $cursor->getNext();
        $this->select += (microtime(true) - $start);
        if(empty($result)){
            ++$this->error;
            if($c['debug']){
                echo "select error\n";
            }
        }
        if($c['debug']){
            print_r($query);
            print_r($result);
            print_r($this->db->{$c['db_name']}->lastError());
        }
    }

    public function getCount(){
        $cursor = $this->collection->find();
        return $cursor->count();
    }
}