<?php

return  [
    'server' => [
        'host' => '0.0.0.0',
        'port' => 9501,
        'settings' => [
            'worker_num' => 1,
            'daemonize' => 0,
            'log_file' => __DIR__ . '/log_file.log'
        ]
    ],
    'tick' => [
        'stat' => [
            [
                'id' => 'ubuntu',
                'host' => '192.168.159.131',
                'port' => 22,
                'username' => 'watcher',
                'password' => '123456',
                'privateKey' => null,
                'publicKey' => null,
                'interval' => 2,
                'stop' => false
            ]
        ]
    ]
];