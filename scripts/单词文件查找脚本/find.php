<?php
/**
 *		find a word from a dir.
 */

echo "input the search dir \n";
$dir = trim(fgets(STDIN));
echo "input the search word \n";
$word = trim(fgets(STDIN));

echo "input extension filter,such as: php c h \n";
$ext = trim(fgets(STDIN));
if($ext){
	$ext = explode(' ',$ext);
}

function get_files($dir,&$files){
	$dir = rtrim($dir,'/');
	$fos = opendir($dir);
	while($file = readdir($fos)){
		if($file == '.' || $file == '..'){
			continue;
		}
		$file = $dir . '/' . $file;
		if(is_dir($file)){
			get_files($file,$files);
		}else{
			$files[] = $file;
		}		
	}
}

get_files($dir,$files);

echo 'file total: ',count($files),"\n";

function is_word_in_file($file,$word,$case=true){
	$content = file_get_contents($file);
	if($case){
		return stripos($content,$word) !== false;
	}else{
		return strpos($content,$word) !== false;
	}
}

foreach($files as $file){
	if($ext && !in_array(pathinfo($file,PATHINFO_EXTENSION),$ext)){
		continue;
	}
	if(is_word_in_file($file,$word)){
		echo $file,"\n";
		$found = true;
	}
}

if(!isset($found)){
	echo "word not found \n";
}
