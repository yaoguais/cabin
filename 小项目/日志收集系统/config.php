<?php

return [
    'SERVER_HOST'           => '0.0.0.0',
    'SERVER_PORT'           => 9505,
    'SERVER_CONFIG'         => [
        'worker_number'  => 1,
        'open_eof_check' => true,
        'package_eof'    => "\r\n"
    ],
    'SERVER_BUFFER_CLASS' => 'fatty\\SwooleBuffer',
    'DB_TYPE'               =>  'mysqli',
    'DB_HOST'               =>  '127.0.0.1',
    'DB_NAME'               =>  'test',
    'DB_USER'               =>  'root',
    'DB_PWD'                =>  '123456',
    'DB_PORT'               =>  '3307',
    'DB_PREFIX'             =>  'pre_',
    'LOG_TYPE'              =>  'File'
];