<?php
/**
 * mysql.php
 *
 * @author    Yong Liu <newtopstdio@163.com>
 * @link      https://github.com/Yaoguais
 * @copyright 2014-2015 TJUT
 * @date      2015-07-07 10:42
 */


//由于当前都是使用的pdo,所以测试mysql也使用pdo进行测试

class BenchMysql extends BenchBase {

    public function __construct($data,$config){
        $this->data = $data;
        $this->config = $config;
        $this->driver = 'mysql';
    }

    public function init(){
        $this->insert = 0;
        $this->select = 0;
        $this->error = 0;
        $c = $this->config;
        $this->db = new PDO($c['dns'],$c['user'],$c['password']);
        if(empty($this->db)){
            $this->error('pdo not enabled');
        }
    }

    public function insert(){
        $c = & $this->config;
        $f = $v = $g = '';
        foreach($this->data as $key=>$val){
            $f .= $g.$key;
            $val .= rand(1000000,9999999);
            $v .= $g."'{$val}'";
            $g = ',';
            $c['key'] = $key;
            $c['val'] = $val;
        }
        $sql = "INSERT INTO {$c['table']}({$f}) VALUES({$v});";
        $start = microtime(true);
        $ret = $this->db->exec($sql);
        $this->insert += (microtime(true) - $start);
        if(empty($ret)){
            ++$this->error;
        }
        if($c['debug']){
            echo $sql,"\n";
            print_r($this->db->errorInfo());
        }
    }

    public function select(){
        $c = $this->config;
        $sql = "SELECT * FROM {$c['table']} WHERE {$c['key']}='{$c['val']}'";
        $start = microtime(true);
        $query = $this->db->query($sql,PDO::FETCH_ASSOC);
        if(empty($query)){
            ++$this->error;
        }else{
            $result = $query->fetch();
        }
        $this->select += (microtime(true) - $start);
        if($c['debug']){
            echo $sql,"\n";
            print_r($result);
        }
    }

    public function getCount(){
        $c = $this->config;
        $sql = "SELECT count(*) as c FROM {$c['table']} limit 1";
        $query = $this->db->query($sql);
        $result = $query->fetch();
        return $result['c'];
    }
}