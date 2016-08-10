<?php
/**
 * BenchBase.php
 *
 * @author    Yong Liu <newtopstdio@163.com>
 * @link      https://github.com/Yaoguais
 * @copyright 2014-2015 TJUT
 * @date      2015-07-07 10:58
 */

abstract class BenchBase {

    protected $insert;
    protected $select;
    protected $error;
    protected $data;
    protected $config;
    protected $db;
    protected $driver = 'none';

    abstract public function init();

    abstract public function insert();

    public function batchInset($number){
        $this->config['insertNumber'] = $number;
        for($i=0;$i<$number;++$i){
            $this->insert();
        }
    }

    abstract public function select();

    public function repeatSelect($number){
        $this->config['selectNumber'] = $number;
        for($i=0;$i<$number;++$i){
            $this->select();
        }
    }

    public function getBenchMark(){
        $c = & $this->config;
        return  $this->driver.":----------------\n".
        "error:  {$this->error} n\n".
        "select({$c['selectNumber']}): {$this->select} s\n".
        "insert({$c['insertNumber']}): {$this->insert} s\n".
        "count: ".$this->getCount()."\n";
    }

    public function error($message){
        exit($message."\n");
    }
}