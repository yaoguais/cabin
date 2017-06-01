<?php
/**
 * Created by PhpStorm.
 * User: yaoguai
 * Date: 15-7-5
 * Time: 下午4:15
 */

$thinkPhpLoad['start'] = microtime(true);
$thinkPhpLoad['memStart'] = memory_get_usage();
if(version_compare(PHP_VERSION,'5.3.0','<'))  die('require PHP > 5.3.0 !');
define('APP_DEBUG',true);
define('RUNTIME_PATH',__DIR__.'/');
define('APP_PATH',__DIR__.'/Application/');
require __DIR__.'/ThinkPHP/ThinkPHP.php';
C(load_config(__DIR__.'/config.php'));
$thinkPhpLoad['end'] = microtime(true);
echo "thinkPHP load:\n",$thinkPhpLoad['end'] - $thinkPhpLoad['start']," s\n";
echo memory_get_usage() - $thinkPhpLoad['memStart']," bytes\n";

require __DIR__.'/bootstrap.php';
(new \fatty\SwooleServer(C('SERVER_HOST'),C('SERVER_PORT'),C('SERVER_CONFIG'),C('SERVER_BUFFER_CLASS')))->start();