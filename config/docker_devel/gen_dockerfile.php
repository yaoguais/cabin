<?php

$dir = __DIR__ . '/dockerfiles';
$tplContent = file_get_contents(__DIR__ . '/dockerfile.tpl');
$fos = opendir($dir);
while ($file = readdir($fos)) {
    if ($file != '.' && $file != '..') {
        $keyword = '### ' . $file . ' ###';
        if (strpos($tplContent, $keyword)) {
            $tplContent = str_replace($keyword, file_get_contents($dir . '/' . $file) . "\n", $tplContent);
            echo $file, "\n";
        }
    }
}
closedir($fos);
file_put_contents(__DIR__ . '/Dockerfile', $tplContent);
