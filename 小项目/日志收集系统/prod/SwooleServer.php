<?php
/**
 * Created by PhpStorm.
 * User: yaoguai
 * Date: 15-7-5
 * Time: 下午4:55
 */

namespace fatty;

class SwooleServer{

    protected $server;
    protected $bufferDriverClass;
    protected $fdBufferMap = [];


    public function __construct($host,$ip,$config,$bufferDriverClass){
        $server = new \swoole_server($host, $ip, SWOOLE_PROCESS, SWOOLE_SOCK_TCP);
        $server->set($config);
        $server->on('receive',array($this,'onReceive'));
        $server->on('close',array($this,'onClose'));
        $this->server = & $server;
        $this->bufferDriverClass = $bufferDriverClass;
    }

    protected function initFdBuffer($fd){
        if(isset($this->fdBufferMap[$fd])){
            return;
        }
        $class = $this->bufferDriverClass;
        $size = C('SERVER_BUFFER_SIZE') ? : 10240;
        $this->fdBufferMap[$fd] = new $class();
        $this->fdBufferMap[$fd]->init($size);
    }

    protected function dispatcherMessage(&$messages){
        //进一步优化是丢到task中处理
        $success = 0;
        foreach($messages as $message){
            if(count($message)>2){
                /*$driver = */array_shift($message);
                $name = array_shift($message);
                D($name)->add($message) && ++$success;
            }
        }
        echo "success: {$success}\n";
    }

    public function onReceive(\swoole_server $server, $fd, $from_id, $message){
        $oef = substr($message,-2);
        if($oef!=="\r\n"){
            $this->initFdBuffer($fd);
            $this->fdBufferMap[$fd]->append($message);
        }else{
            if(isset($this->fdBufferMap[$fd])){
                $this->fdBufferMap[$fd]->append($message);
                $message = $this->fdBufferMap[$fd]->get();
                $this->fdBufferMap[$fd]->clear();
            }
            $result = Package::parsePackage($message);
            if($result['success']>0){
                $this->dispatcherMessage($result['messages']);
            }else if(!$result['ok']){
                echo "parse error:\n";
                print_r($result['error']);
            }
        }
    }

    public function onClose(\swoole_server $server, $fd,$from_id){
        if(isset($this->fdBufferMap[$fd])){
            $this->fdBufferMap[$fd]->destroy();
            $this->fdBufferMap[$fd] = null;
        }
    }

    public function start(){
        $this->server->start();
    }
}