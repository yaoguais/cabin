<?php

//文档：http://www.cnblogs.com/xiaowu/archive/2012/09/18/2690677.html
/**
 * 向服务器中的log表中添加日志数据
 */

if($argc<3){
    exit("php client.php 10 100 tag\n");
}
$tag = $argv[3] ? : "result";
$sock = socket_create(AF_INET, SOCK_DGRAM, SOL_UDP);
$data = ['log','data'=>'a message'];
$data = serialize($data);
$len = strlen($data);
$loop = $originLoop = $argv[1] ? : 10;
$number = $argv[2] ? : 100;
$total = $originLoop*$number;
$errorNumber = 0;
$timeList = [];
$secondUnit = 1000000;
$totalSecond = 0;
while(--$loop>=0){
    for($i=0;$i<$number;++$i){
        $start = microtime(true);
        $send = socket_sendto($sock, $data, $len, 0, '127.0.0.1', 9905);
        if($send){
            $timeList[$loop][$i] = (microtime(true) - $start)*$secondUnit;
            $totalSecond += $timeList[$loop][$i];
        }else{
            ++$errorNumber;
        }

        //echo "{$timeList[$loop][$i]}\n";
    }
}
socket_close($sock);

echo "\n\n\n\nresult:\nerror:{$errorNumber} loop:{$originLoop} number:{$number}\n";
$loopTimeList = [];
foreach($timeList as $k=>$numberList){
    $loopTimeList[$k] = array_sum($numberList)/$number;
    echo "{$loopTimeList[$k]}\n";
}
echo "finally: ".(array_sum($loopTimeList)/$originLoop),"/{$secondUnit}s\n";


$lockFile = 'receive-count.txt';
sleep(120);
$sendNumber = $total-$errorNumber;
$result = "send:\nnumber:$total\n";
$result .= "error:$errorNumber\nsuccess:$sendNumber\n";
$result .= "takes:{$totalSecond}/{$secondUnit}s\n";
$result .= "rqs:".($total/$totalSecond*$secondUnit)." n/s\n";
$success = intval(file_get_contents($lockFile));
$result .= "receive:\nsuccess:$success\n";
$result .= "last:".(($sendNumber - $success)/$sendNumber*100)."%\n\n\n";
file_put_contents("{$tag}.txt",$result,FILE_APPEND);
