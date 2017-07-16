<?php
$file = __DIR__ . '/mysqlnd.trace.2';
$dir = __DIR__ . '/mysqlnd.trace.2.separate';
mkdir($dir);
$contents = file_get_contents($file);
$rows = explode("\n", $contents);
foreach ($rows as $row) {
    $items = explode(':', $row, 2);
    if (strlen($items[0]) > 0) {
        $outFile = $dir . '/' . $items[0] . '.log';
        file_put_contents($outFile, $row . "\n", FILE_APPEND);
    }
}
