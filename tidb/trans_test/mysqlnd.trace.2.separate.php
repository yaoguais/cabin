<?php
$file = __DIR__ . '/mysqlnd.trace.2';
$dir = __DIR__ . '/mysqlnd.trace.2.separate';
mkdir($dir);
$fp = fopen($file, "r");
while($line = fgets($fp)) {
    $items = explode(':', $line, 2);
    if (strlen($items[0]) > 0) {
        $outFile = $dir . '/' . $items[0] . '.log';
        file_put_contents($outFile, $line, FILE_APPEND);
    }
}
fclose($fp);
