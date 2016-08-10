<?php

require_once __DIR__ . '/watcher.php';
$config = require __DIR__ . '/config.php';
use \watcher\Server;
use \watcher\TopTick;
use \watcher\DefaultServerHandler;
use \watcher\TickManager;
use \watcher\TopProtocol;
use \watcher\ProtocolManager;
$server = Server::getInstance();
$server->host = $config['server']['host'];
$server->port = $config['server']['port'];
$server->settings = $config['server']['settings'];
$tickManager = TickManager::getInstance();
foreach($config['tick']['stat'] as $statConfig){
    $tickManager->addTick(new TopTick($statConfig));
}
$protocolManager = ProtocolManager::getInstance();
$protocolManager->addProtocol(new TopProtocol());
$server->addHandle(new DefaultServerHandler());
$server->run();
