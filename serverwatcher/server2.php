<?php

require_once __DIR__ . '/watcher.php';
require_once __DIR__ . '/request.php';
require_once __DIR__ . '/disk.php';
$config = require __DIR__ . '/config.php';
use \watcher\Server;
use \watcher\DefaultServerHandler;
use \watcher\TickManager;
use \watcher\TopProtocol;
use \watcher\RequestProtocol;
use \watcher\DiskProtocol;
use \watcher\ProtocolManager;
$server = Server::getInstance();
$server->host = $config['server']['host'];
$server->port = $config['server']['port'];
$server->settings = $config['server']['settings'];
$tickManager = TickManager::getInstance();
foreach($config['tick'] as $tickConfig){
    $class = $tickConfig['class'];
    $tickManager->addTick(new $class($tickConfig));
}
$protocolManager = ProtocolManager::getInstance();
$protocolManager->addProtocol(new TopProtocol());
$protocolManager->addProtocol(new RequestProtocol());
$protocolManager->addProtocol(new DiskProtocol());
$server->addHandle(new DefaultServerHandler());
$server->run();