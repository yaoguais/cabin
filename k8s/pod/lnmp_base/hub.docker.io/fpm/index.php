<?php


foreach([$_POST, $_GET, $_SERVER, $_COOKIE] as $g) {
	echo json_encode($g), "\n";
}

echo "ok\n";

