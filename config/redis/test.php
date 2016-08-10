<?php

if($argc != 2){
	exit("php test.php basedir\n");
}

$basedir = rtrim($argv[1],'/').'/';

function get_dirs_by_dir($dir){
	$dirs = [];
	$fos = opendir($dir);
	while($name = readdir($fos)){
		if($name == '.' || $name == '..'){
			continue;
		}
		if(is_dir($dir.$name)){
			$dirs[] = $dir.$name;
		}
	}
	closedir($fos);
	return $dirs;
}

$dirs = get_dirs_by_dir($basedir);

//print_r($dirs);

function array_repeat($arr)
{
	if(!is_array($arr)) return null;

	$arr1 = array_count_values($arr);

	$newArr = array();

	foreach($arr1 as $k=>$v)
	{
		if($v>1) array_push($newArr,$k);
	}
	return $newArr;
}


function check_config_file($dirs){
	$ports = [];
	foreach($dirs as $dir){
		$config = $dir.'/redis.conf';
		$content = file_get_contents($config);
		
		$dirname = substr($dir,strrpos($dir,'/')+1);
		$regex = preg_replace("/\d/",'\d',$dirname);
		//echo $regex,"\n";
		preg_match_all("/$regex/is",$content,$m);
		if(!isset($m[0])){
			echo $config," no dirname error\n";
		}else{
			if(count($m[0]) != 5){
				echo $config," same dirname must be 5 ,now is ",count($m[0]),"\n";
			}
			foreach($m[0] as $match){
				if($match != $dirname){
					echo $config," filename error\n";
				}
			}
		}
		if(strpos($dirname,'slave') !== false){
			if(preg_match("/#\s*?slaveof/is",$content)){
				echo $config," slaveof error\n";
			}
		}
		preg_match("/port\s+(\d+)/is",$content,$m);
		if(!isset($m[1])){
			echo $config," port not set error\n";
		}else{
			if(in_array($m[1],$ports)){
				echo $config," found same port error\n";
				echo $dirs[array_search($m[1],$ports)]," is the target\n";
			}
			$ports[] = $m[1];
		}
	}
	
}

check_config_file($dirs);

echo "-------------------------------------------------------------------------------\n";
