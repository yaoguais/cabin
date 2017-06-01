<?php

//文档：http://www.cnblogs.com/xiaowu/archive/2012/09/18/2690677.html
/**
 * 向服务器中的log表中添加日志数据
 */
define('CHECK_CHAR','0');

if($argc<3){
    exit("php client.php 10 100 tag\n");
}
$loop = $originLoop = $argv[1] ? : 10;
$number = $argv[2] ? : 100;
$tag = $argv[3] ? : "result";
$totalNumber = $originLoop*$number;


$start = microtime(true);
$sock = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
$con=socket_connect($sock,'127.0.0.1',9905);
if(!$con){socket_close($sock);exit("connect error\n");}

$data = ['log','data'=>'a message'];
$data = serialize($data);
$len = strlen($data);
$packLenString = pack('N',$len);
$out = '';
while(--$loop>=0){
    for($i=0;$i<$number;++$i){
        $out .= $packLenString.CHECK_CHAR.$data;
    }
}
$out .= "\r\n";
$len = strlen($out);

$writeLen = socket_write($sock,$out,$len);
socket_shutdown($sock);
socket_close($sock);

$result = "data number:{$totalNumber}\ndata length:{$len}\nwrite length:{$writeLen}\ntake:".(microtime(true)-$start)."s\nsuccess record:\n\n";
echo $result;
file_put_contents("{$tag}.txt",$result,FILE_APPEND);
