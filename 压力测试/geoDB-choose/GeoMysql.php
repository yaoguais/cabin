<?php
/**
 * mysql.php
 *
 * @author    Yong Liu <newtopstdio@163.com>
 * @link      https://github.com/Yaoguais
 * @copyright 2014-2015 TJUT
 * @date      2015-07-07 18:18
 */


/*
 * create table pre_geo(
    id int(11) not null auto_increment primary key,
    geo geometry not null,
    SPATIAL KEY(geo)
   );
    SELECT id,ST_Distance_Sphere(Point(160,30), geo) as distance_in_meters from pre_geo order by distance_in_meters desc limit 5
 */
class GeoMysql extends GeoBase{

    public function __construct($config){
        $this->config = $config;
        $this->driver = 'mysql';
        $this->init();
    }

    public function init(){
        $c = $this->config;
        $this->db = new PDO($c['dns'],$c['user'],$c['password']);
        if(empty($this->db)){
            $this->error('pdo not enabled');
        }
    }

    public function repeatSelect($number){
        $this->selectTime = $success = 0;
        for($i=0;$i<$number;++$i){
            $lat = rand(-180,179)+1/rand(2,99);
            $lon = rand(-90,89)  + 1/rand(2,99);
            $sql = "SELECT id,ST_Distance_Sphere(Point($lat,$lon), geo) as distance_in_meters from pre_geo order by distance_in_meters desc limit 5";
            $start = microtime(true);
            $query = $this->db->query($sql);
            if($query){
                $this->selectTime += (microtime(true) - $start);
                $result = $query->fetchAll(PDO::FETCH_ASSOC);
                if($result){
                    $this->record($result);
                    ++$success;
                }else{
                    echo "failed\n";
                }
            }else{
                $this->selectTime += (microtime(true) - $start);
                echo "failed\n";
            }
        }
        $msg = "number:$number\nsuccess:$success\ntime:{$this->selectTime}\n\n";
        $this->record($msg,'result');
        return $this->selectTime;
    }

    public function record($msg,$file='mysql'){
        if(is_array($msg)){
            $msg = var_export($msg,true);
        }
        file_put_contents(__DIR__.'/'.$file.'.txt',$msg,FILE_APPEND);
    }

    public function createTestData($number=100000){
        $success = 0;
        for($i=0;$i<$number;++$i){
            $lat = rand(-180,179)+1/rand(2,99);
            $lon = rand(-90,89)  + 1/rand(2,99);
            $sql = "INSERT INTO pre_geo(geo) VALUES(Point($lat,$lon));";
            $this->db->exec($sql) && ++$success;
        }
        echo "success create: $success \n";
    }
}