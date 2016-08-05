<?php
/**
 * GeoBase.php
 *
 * @author    Yong Liu <newtopstdio@163.com>
 * @link      https://github.com/Yaoguais
 * @copyright 2014-2015 TJUT
 * @date      2015-07-07 18:19
 */

abstract class GeoBase {

    protected $db;
    protected $driver;
    protected $config;
    protected $selectTime = 0;

    abstract function init();

    public function error($message){
        exit($message."\n");
    }
}