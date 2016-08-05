<?php
/**
 * Created by PhpStorm.
 * User: yaoguai
 * Date: 15-7-5
 * Time: 下午5:36
 */

define('FATTY_ROOT',__DIR__.'/prod/');
spl_autoload_register(function($class){
    static $classes = [];
    if(!isset($classes[$class])){
        if(substr($class,0,6) == 'fatty\\'){
            include FATTY_ROOT.substr($class,6).'.php';
        }
    }
});