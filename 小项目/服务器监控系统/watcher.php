<?php

namespace watcher {

    class Server extends Listener{

        protected $server;
        protected $events = ['workerStart'=>true, 'open'=>true, 'message'=>true, 'close'=>true];

        public $host = '0.0.0.0';
        public $port = 9501;
        public $settings;

        protected static $instance;

        private function __construct(){}

        public static function getInstance(){

            if(!self::$instance){
                self::$instance = new self;
            }

            return self::$instance;
        }

        public function run(){

            if(!$this->host){
                throw new ServerException('host can not be empty',ServerException::HOST_EMPTY);
            }
            if(!$this->port){
                throw new ServerException('port can not be empty',ServerException::PORT_EMPTY);
            }
            $server = new \swoole_websocket_server($this->host, $this->port);
            $this->settings['worker_num'] = 1;
            $server->set($this->settings);
            $server->on('workerStart',[$this,'workerStart']);
            $server->on('open',[$this,'open']);
            $server->on('message',[$this,'message']);
            $server->on('close',[$this,'close']);
            $this->server = $server;
            $this->server->start();
        }

        protected function workerStart($server, $workerId){

        }

        protected function open($server, $request){

            $fd = $request->fd;
            $user = new User();
            $user->fd = $fd;
            $user->ip = $this->getIpByFd($fd);
            UserManager::getInstance()->addUser($user);
        }

        protected function message($server, $frame){

        }

        protected function close($server, $fd){

            UserManager::getInstance()->removeUserByFd($fd);
        }

        public function sendMessage($fd,$message){

            if($this->server){
                return $this->server->push($fd,$message);
            }

            return false;
        }

        public function getIpByFd($fd){

            if($this->server){
                $info = $this->server->connection_info($fd);
                if(isset($info['remote_ip'])){
                    return $info['remote_ip'];
                }
            }

            return false;
        }
    }

    class ServerException extends \Exception {

        const HOST_EMPTY = 1;
        const PORT_EMPTY = 2;
    }

    interface HandleInterface{}

    abstract class ServerHandle implements  HandleInterface{

        public abstract function beforeWorkerStart($server, $workerId);
        public abstract function afterWorkerStart($server, $workerId, $return = null);
        public abstract function beforeOpen($server, $request);
        public abstract function afterOpen($server, $request, $return = null);
        public abstract function beforeMessage($server, $frame);
        public abstract function afterMessage($server, $frame, $return = null);
        public abstract function beforeClose($server, $fd);
        public abstract function afterClose($server, $fd, $return = null);
    }

    class DefaultServerHandler extends ServerHandle {

        protected $initTick = false;

        protected function getUser($server, $fd){

            if($user = UserManager::getInstance()->getUserByFd($fd)){
                return $user;
            }
            $server->close($fd);

            return null;
        }

        public function beforeWorkerStart($server, $workerId){

            if (!$server->taskworker) {
                if(!$this->initTick) {
                    $this->initTick = true;
                    $server->tick(TickInterface::INTERVAL, [TickManager::getInstance(), 'tick'], $server);
                }
            }
        }

        public function afterWorkerStart($server, $workerId, $return = null){

        }

        public function beforeOpen($server, $request){

        }

        public function afterOpen($server, $request, $return = null){

            if($user = $this->getUser($server, $request->fd)){
                ProtocolManager::getInstance()->init($user);
            }
        }

        public function beforeMessage($server, $frame){

            if($user = $this->getUser($server, $frame->fd)){
                $packet = new Packet();
                if($packet->unSerializeFromRaw($frame->data)){
                    ProtocolManager::getInstance()->dispatcher($user,$packet);
                }
            }
        }

        public function afterMessage($server, $frame, $return = null){

        }

        public function beforeClose($server, $fd){

            if($user = $this->getUser($server, $fd)){
                ProtocolManager::getInstance()->clear($user);
            }
        }

        public function afterClose($server, $fd, $return = null){

        }
    }

    class Packet{

        const COMMAND_UNKNOWN = 0;
        const COMMAND_FLOW     = 105;

        const ERROR_NOT_EXISTS = 0;
        const ERROR_COMMAND = 1;
        const ERROR_DATA    = 2;

        public $command = self::COMMAND_UNKNOWN;
        public $code    = self::ERROR_NOT_EXISTS;
        public $message;
        public $data;
        protected $commands;

        public function serialize(){

            return json_encode([
                'command' => $this->command,
                'code' => $this->code,
                'message' => $this->message,
                'data' => $this->data
            ]);
        }

        public function unSerializeFromRaw($raw){

            $data = json_decode($raw,true);
            if($data && isset($data['command'])){
                $this->command = intval($data['command']);
                isset($data['code']) && ($this->code = intval($data['code']));
                isset($data['message']) && ($this->message = $data['message']);
                isset($data['data']) && ($this->data = $data['data']);
                return true;
            }

            return false;
        }
    }

    class TopPacket extends Packet{

        const COMMAND_GET       = 100;
        const COMMAND_PUSH      = 101;
        const COMMAND_AUTH      = 102;
        const COMMAND_NEED_AUTH = 103;
        const COMMAND_RE_FLOW   = 104;

        const ERROR_USERNAME = 100;
        const ERROR_PASSWORD = 101;
        const ERROR_AUTH     = 102;

        protected $commands = [
            self::COMMAND_GET => true,
            self::COMMAND_AUTH => true,
            self::COMMAND_RE_FLOW => true
        ];
    }

    interface ProtocolInterface{

        public function init(User $user);
        public function handle(User $user, Packet $packet);
        public function clear(User $user);
    }

    class Protocol implements ProtocolInterface{

        public function init(User $user){}
        public function handle(User $user, Packet $packet){}
        public function clear(User $user){}
    }

    class ProtocolManager {

        /* @var Protocol[] */
        protected $protocols;

        protected static $instance;

        private function __construct(){}

        public static function getInstance(){

            if(!self::$instance){
                self::$instance = new self;
            }

            return self::$instance;
        }

        public function addProtocol(Protocol $protocol){

            $this->protocols[get_class($protocol)] = $protocol;
        }

        public function removeProtocol(Protocol $protocol){

            unset($this->protocols[get_class($protocol)]);
        }

        public function init(User $user){

            foreach($this->protocols as $protocol){
                $protocol->init($user);
            }
        }

        public function dispatcher(User $user, Packet $packet){

            foreach($this->protocols as $protocol){
                $protocol->handle($user, clone $packet);
            }
        }

        public function clear(User $user){

            foreach($this->protocols as $protocol){
                $protocol->init($user);
            }
        }
    }

    abstract class TickProtocol extends Protocol{

        abstract protected function parseStatDataToSimpleArray(DataBase $data);

        public function handleFlowStatData($tick, $statData){

            $userManager =  UserManager::getInstance();
            $userManager->each(function($user) use($userManager, $tick, $statData){
                /* @var TickBase $tick */
                $id = $tick->getId();
                $type = $tick->getType();
                if($user->auth && (!isset($user->data[$type][$id]['immediate']) || false !== $user->data[$type][$id]['immediate'])){
                    $message = $this->getFlowMessage($tick, $statData);
                    $userManager->sendMessage($user,$message);
                }
            });
        }

        protected function getFlowMessage(TickBase $tick, $statData){

            $packet = new Packet();
            $packet->command = Packet::COMMAND_FLOW;
            $packet->data[$tick->getId()] = $this->getDataForId($tick->getType(),[$this->parseStatDataToSimpleArray($statData)], true);
            return $packet->serialize();
        }

        protected function getDataForId($type, $list, $flow, $start = -1, $end = -1){

            return ['type' => $type, 'start' => $start, 'end' => $end, 'flow' => $flow ? 1 : 0, 'list' => $list];
        }

        protected function getList($id, $modelName, &$start, &$end){

            if(!$id){
                return null;
            }
            $defaultTime = 300;
            $defaultMax  = 50;
            $now = time();
            $start = $start ? : $now - $defaultTime;
            $end = $end ? : $now;
            $list = [];
            if($id){
                $statDataList = DataManager::get($id, $modelName , $start, $end);
                //default every id return 500 max
                $count = count($statDataList);
                if($count > $defaultMax){
                    $step = floor($count/ $defaultMax);
                    $index = 0;
                    foreach($statDataList as $startData){
                        if($index == 0){
                            $list[] = $this->parseStatDataToSimpleArray($startData);
                        }
                        ++$index;
                        if($index > $step){
                            $index = 0;
                        }
                    }
                }else{
                    foreach($statDataList as $startData){
                        $list[] = $this->parseStatDataToSimpleArray($startData);
                    }
                }
            }
            return $list;
        }
    }

    class TopProtocol extends TickProtocol{

        public function init(User $user){

            $packet = new TopPacket();
            $packet->command = TopPacket::COMMAND_NEED_AUTH;
            UserManager::getInstance()->sendMessage($user,$packet->serialize());
        }

        public function handle(User $user, Packet $packet){

            $userManager = UserManager::getInstance();
            $tickManager = TickManager::getInstance();
            switch($packet->command){
                case TopPacket::COMMAND_UNKNOWN:
                    return null;
                case TopPacket::COMMAND_AUTH:
                    if($user->auth){

                    }else if(!isset($packet->data['username'])){
                        $packet->code = TopPacket::ERROR_USERNAME;
                    }else if(!isset($packet->data['password'])){
                        $packet->code = TopPacket::ERROR_PASSWORD;
                    }else{
                        $user->username = $packet->data['username'];
                        $user->password = $packet->data['password'];
                        $tickManager->each(function($tick) use($user){
                            if($tick instanceof TopTick){
                                $tick->auth($user);
                            }
                        });
                        if(!$user->auth){
                            $packet->code = TopPacket::ERROR_AUTH;
                        }
                    }
                    $packet->data = null;
                    $userManager->sendMessage($user, $packet->serialize());
                    break;
                case TopPacket::COMMAND_GET:
                    if(!$user->auth){
                        $packet->code = TopPacket::ERROR_AUTH;
                    }else{
                        $list = [];
                        if(!$packet->data){
                            $idList = $tickManager->each(function($tick){
                                if($tick instanceof TopTick){
                                    return $tick->getId();
                                }
                            });
                            if(is_array($idList)){
                                $start = $end = null;
                                foreach($idList as $id){
                                    if($id && ($data = $this->getList($id, 'watcher\\TopData',$start,$end))){
                                        $tick = $tickManager->getTickById($id);
                                        $type = $tick->getType();
                                        $user->data[$type][$id]['immediate'] = false;
                                        $list[$id] = $this->getDataForId($type, $data, false, $start, $end);
                                    }
                                }
                            }
                        }
                        if($list){
                            $packet->data = $list;
                            $userManager->sendMessage($user, $packet->serialize());
                        }
                    }
                    break;
                case TopPacket::COMMAND_RE_FLOW:
                    if(!$packet->data){
                        unset($user->data['stat']);
                    }else{
                        $user->data['stat'] = null;
                        if(is_array($packet->data)){
                            foreach($packet->data as $id){
                                $user->data['stat'][$id]['immediate'] = true;
                            }
                        }
                    }
                    $packet->data = null;
                    $userManager->sendMessage($user, $packet->serialize());
                    break;
                default:
                    return null;
            }
        }

        protected function parseStatDataToSimpleArray(DataBase $model){

            return [
                $model->timestamp,
                $model->cpuUser,
                $model->cpuSystem,
                $model->memory,
                $model->loadAverage,
                $model->swap,
                $model->buffers,
                $model->cached
            ];
        }
    }

    class TickException extends \Exception{

        const INTERVAL_ERROR = 1;
        const ID_ERROR       = 2;
        const TYPE_ERROR     = 3;
    }

    interface TickInterface {

        const INTERVAL = 1000;
        public function tick($id, $server);
    }

    abstract class TickBase implements TickInterface{

        protected $interval;
        protected $id;
        protected $type;
        protected $stop;

        public function __construct($config){

            if(!isset($config['id']) || null === $config['id'] || '' === $config['id']){
                throw new TickException('id can not be null or empty string', TickException::ID_ERROR);
            }
            $this->setId($config['id']);
            if(!is_int($config['interval']) || $config['interval'] < 1){
                throw new TickException('interval must be integer and >=1', TickException::INTERVAL_ERROR);
            }
            $this->setInterval($config['interval']);
            if(!isset($config['type'])){
                throw new TickException('type can not be null or empty string', TickException::TYPE_ERROR);
            }
            $this->setType($config['type']);
            $this->stop = isset($config['stop']) ? $config['stop'] : false;
        }

        public function setId($id){

            if(null === $this->id){
                $this->id = $id;
            }
        }

        public function getId(){

            return $this->id;
        }

        public function setType($type){

            $this->type = $type;
        }

        public function getType(){

            return $this->type;
        }

        public function setInterval($interval){

            $this->interval = $interval < 1 ? 1 : intval($interval);
        }

        public function getInterval(){

            return $this->interval;
        }

        public function setStop($bool){

            $this->stop = $bool;
        }

        public function isStop(){

            return $this->stop;
        }
    }

    class TickManager implements TickInterface{

        protected $ticks;
        protected $timestamp;

        protected static $instance;

        private function __construct(){}

        public static function getInstance(){

            if(!self::$instance){
                self::$instance = new self;
            }

            return self::$instance;
        }

        public function addTick(TickBase $tick){

            if(null === $this->getTickById($tick->getId())){
                $interval = $tick->getInterval();
                $this->ticks[$interval][] = $tick;
            }
        }

        public function removeTick(TickBase $tick){

            if($this->ticks){
                $interval = $tick->getInterval();
                if(isset($this->ticks[$interval])){
                    foreach($this->ticks[$interval] as $k=>$object){
                        if($object === $tick){
                            unset($this->ticks[$interval][$k]);
                        }
                    }
                }
            }
        }

        public function getTickById($id){

            if($this->ticks){
                foreach($this->ticks as $ticks){
                    foreach($ticks as $tick){
                        if($tick->getId() === $id){
                            return $tick;
                        }
                    }
                }
            }

            return null;
        }

        public function each($callback){

            $list = null;
            if($this->ticks){
                foreach($this->ticks as $ticks){
                    foreach($ticks as $tick){
                        $ret = $callback($tick);
                        if( $ret === false){
                            return false;
                        }
                        $list[] = $ret;
                    }
                }
            }

            return $list;
        }

        public function removeTickById($id){

            if($this->ticks){
                foreach($this->ticks as &$ticks){
                    foreach($ticks as $k=>$tick){
                        if($tick->getId() === $id){
                            unset($ticks[$k]);
                        }
                    }
                }
            }
        }

        public function tick($id, $server){

            $now = time();
            if(!$this->timestamp){
                $this->timestamp = $now;
                $difference = 0;
            }else{
                $difference = $now - $this->timestamp;
            }
            if($this->ticks){
                foreach($this->ticks as $interval => $ticks){
                    if($difference % $interval == 0){
                        foreach($ticks as $tick){
                            if(!$tick->isStop()){
                                if($tick->getId()){
                                    $tick->tick($id, $server);
                                }else{
                                    Logger::log('tick must call parent::__construct', Logger::LEVEL_ERROR);
                                }
                            }
                        }
                    }
                }
            }
        }
    }

    class TopTick extends TickBase{

        protected $sshConfig;

        public function __construct($config){

            parent::__construct($config);
            SshManager::getInstance()->addConnection($config);
            $this->sshConfig = $config;
        }

        public function tick($id, $server){

            /* @var Ssh $ssh */
            $ssh = SshManager::getInstance()->getConnection($this->sshConfig);
            if($ret = $ssh->exec('top -d 0.01 -n 2')){
                $data = new TopData();
                if($data->unSerializeFromTopOutput($ret)){
                    DataManager::save($this->getId(), $data);
                    (new TopProtocol())->handleFlowStatData($this, $data);
                }else{
                    Logger::log('parse top output error, ticker ' . $this->getId(), Logger::LEVEL_ERROR);
                }
            }else{
                Logger::log('ssh exec error, '. $ssh->getLastError(), Logger::LEVEL_ERROR);
                $this->setStop(true);
            }
        }

        public function setStop($bool){

            $this->stop = $bool;
            if($bool){
                SshManager::getInstance()->removeConnection($this->sshConfig);
            }
        }

        public function auth(User $user){

            $user->auth = true;
            return true;
        }
    }

    class User{

        public $fd;
        public $username;
        public $password;
        public $ip;
        public $auth;
        public $data;

        public static function checkPassword($input,$password){

            return $input && $input == md5($password);
        }
    }

    class UserManager {

        protected $users;
        protected static $instance;

        public $multiLogin;

        private function __construct(){}

        public static function getInstance(){

            if(!self::$instance){
                self::$instance = new self;
            }

            return self::$instance;
        }

        public function addUser(User $user){

            $fd = $user->fd;
            if($fd <= 0 ){
                throw new UserException('fd can be empty', UserException::FD_ERROR);
            }
            if(isset($this->users[$fd])){
                throw new UserException('fd already exist', UserException::FD_EXISTS);
            }
            if(!$this->checkMultiLogin($user)){
                return false;
            }
            $this->users[$fd] = $user;

            return true;
        }

        public function getUserByFd($fd){

            if($fd <= 0 ){
                throw new UserException('fd can be empty', UserException::FD_ERROR);
            }
            if($this->users){
                foreach($this->users as $user){
                    if($user->fd == $fd){
                        return $user;
                    }
                }
            }
        }

        public function removeUserByFd($fd){

            if($fd <= 0 ){
                throw new UserException('fd can be empty', UserException::FD_ERROR);
            }
            if($this->users){
                foreach($this->users as $k=>$user){
                    if($user->fd == $fd){
                        unset($this->users[$k]);
                    }
                }
            }
        }
        public function each($callback){

            $list = [];
            if($this->users){
                foreach($this->users as $user){
                    $ret = $callback($user);
                    if($ret === false){
                        return false;
                    }
                    $list[] = $ret;
                }
            }

            return $list;
        }

        public function checkMultiLogin(User $user){

            if(!$this->multiLogin && $this->users){
                foreach($this->users as $object){
                    if($object->username === $user->username){
                        return false;
                    }
                }
            }

            return true;
        }

        public function sendMessage(User $user, $message){

            return Server::getInstance()->sendMessage($user->fd, $message);
        }
    }

    class UserException extends \Exception{

        const FD_ERROR = 1;
        const FD_EXISTS = 2;
    }

    class Logger {

        const LEVEL_TRACE='trace';
        const LEVEL_WARNING='warning';
        const LEVEL_ERROR='error';
        const LEVEL_INFO='info';
        const LEVEL_PROFILE='profile';

        public static function log($msg,$level = self::LEVEL_ERROR){

            echo $msg, "\n";//save message to swoole log file
        }
    }

    class Listener {

        protected $handles;
        protected $events;

        public function isHandleExists($handle){

            if( !$handle || !( $handle instanceof HandleInterface ) ){
                throw new ListenerException('listener must instance of Handle',ListenerException::LISTENER_INSTANCE_ERROR);
            }

            return isset( $this->handles[get_class($handle)] );
        }

        public function addHandle($handle){

            if(!$this->isHandleExists($handle)){
                $this->handles[get_class($handle)] = $handle;
            }
        }

        public function removeHandle($handle){

            if($this->isHandleExists($handle)){
                unset($this->handles[get_class($handle)]);
            }
        }

        public function __call($name, $arguments){

            if(!method_exists($this,$name)){
                throw new ListenerException('method not exits',ListenerException::METHOD_NOT_EXISTS);
            }
            if($this->events && isset($this->events[$name])){
                if($this->handles){
                    foreach($this->handles as $handle){
                        $callback = 'before'.ucwords($name);
                        if((null !== ($ret = call_user_func_array([$handle,$callback],$arguments)))){
                            return $ret;
                        }
                    }
                }
                $ret = call_user_func_array([$this,$name],$arguments);
                if($this->handles){
                    $arguments[] = $ret;
                    foreach($this->handles as $handle){
                        $callback = 'after'.ucwords($name);
                        if((null !== ($ret = call_user_func_array([$handle,$callback],$arguments)))){
                            return $ret;
                        }
                    }
                }
            }else{
                $ret = call_user_func([$this,$name],$arguments);
            }

            return $ret;
        }
    }

    class ListenerException extends \Exception{

        const LISTENER_INSTANCE_ERROR = 1;
        const METHOD_NOT_EXISTS       = 2;
    }

    class SshException extends \Exception{

        const HOST_EMPTY       = 1;
        const USERNAME_EMPTY   = 2;
        const VERITY_EMPTY     = 3;
        const CONNECT_ERROR    = 4;
        const METHOD_NOT_EXITS = 5;
        const FORGET_TO_ADD    =6;
    }

    class SshManager {

        protected $connections;
        protected $referenceCount;
        protected static $instance;

        private function __construct(){}

        public static function getInstance(){

            if(!self::$instance){
                self::$instance = new self;
            }

            return self::$instance;
        }

        public function addConnection($config){

            $id = $this->getIdByConfig($config);
            if(isset($this->referenceCount[$id]) && $this->referenceCount[$id] > 0){
                ++$this->referenceCount[$id];
            }else{
                $this->referenceCount[$id] = 1;
                $ssh = new Ssh();
                foreach($config as $attr=>$value){
                    if(property_exists($ssh, $attr)){
                        $ssh->$attr = $value;
                    }
                }
                $this->connections[$id] = $ssh;
            }
        }

        public function getConnection($config){

            $id = $this->getIdByConfig($config);
            if(isset($this->referenceCount[$id]) && $this->referenceCount[$id] > 0){
                return $this->connections[$id];
            }else{
                throw new SshException('forget to add connection, get must after add', SshException::FORGET_TO_ADD);
            }
        }

        public function removeConnection($config){

            $id = $this->getIdByConfig($config);
            if(isset($this->referenceCount[$id])){
                if($this->referenceCount[$id] > 1){
                    --$this->referenceCount[$id];
                }else{
                    $this->referenceCount[$id] = 0;
                    $this->connections[$id]->close();
                    unset($this->connections[$id]);
                }
            }
        }

        protected function getIdByConfig($config){

            $id = '';
            if(isset($config['host'])){
                $id .= $config['host'];
            }
            $id .= '_';
            if(isset($config['username'])){
                $id .= $config['username'];
            }

            return $id;
        }
    }

    class Ssh {

        const VERSION = 2;

        protected $conn;
        protected $lastError;
        protected $auth;

        public $host;
        public $port = 22;
        public $username;
        public $password;
        public $privateKey;
        public $publicKey;
        public $phrase;
        public $connMethods;
        public $connCallback;


        protected function check(){

            if(!$this->host){
                throw new SshException('host can not be empty', SshException::HOST_EMPTY);
            }
            if(!$this->username){
                throw new SshException('username can not be empty', SshException::USERNAME_EMPTY);
            }
            if(!$this->password && !$this->privateKey){
                throw new SshException('private key or password can not be both empty', SshException::VERITY_EMPTY);
            }
        }

        public function connect(){

            $this->check();
            if(!$this->conn){
                if(!isset($this->connMethods['hostkey'])){
                    $this->connMethods['hostkey'] = 'ssh-rsa';
                }
                $this->conn = ssh2_connect($this->host, $this->port, $this->connMethods, $this->connCallback);
            }
            if(!$this->conn){
                $this->lastError = 'connection failed';
                return false;
            }else{
                return true;
            }
        }

        public function close(){

            unset($this->conn);
            $this->conn = null;
        }

        public function auth(){

            if(!$this->conn && !$this->connect()){
                return false;
            }
            if($this->privateKey){
                $ret = ssh2_auth_pubkey_file($this->conn, $this->username, $this->publicKey, $this->privateKey, $this->phrase);
            }else{
                $ret = ssh2_auth_password($this->conn, $this->username, $this->password);
            }
            if(!$ret){
                $this->lastError = 'auth failed';
                return false;
            }else{
                return true;
            }
        }

        public function getLastError(){

            return $this->lastError;
        }

        public function exec($command, $output = true, $reTry = 1, $pty = null, $env = null , $width = null, $height = null, $width_height_type = null) {

            $pty = $pty ? : 'ansi';
            while($reTry >= 0){
                --$reTry;
                if(!$this->auth && !$this->auth()){
                    return false;
                }
                $this->auth = true;
                $stream = ssh2_exec($this->conn, $command, $pty, $env, $width, $height, $width_height_type);
                if(false === $stream){
                    $this->lastError = 'shell execute failed';
                    $this->auth = false;
                    continue;
                }
                if(!$output){
                    return true;
                }
                stream_set_blocking($stream, true);
                $content = stream_get_contents($stream);
                fclose($stream);
                return $content;
            }

            return false;
        }

        public function __call($name, $arguments){

            $method = 'ssh2_' . $name;
            if(!method_exists($this,$method)){
                throw new SshException('method not exists', SshException::METHOD_NOT_EXITS);
            }

            return call_user_func($this->conn, $arguments);
        }
    }

    abstract class DataBase{

        public $timestamp = -1;
        abstract  public function serializeToRaw();
        abstract  public function unSerializeFromRaw($raw);
    }

    class TopData extends  DataBase{

        public $cpuUser = -1;
        public $cpuSystem = -1;
        public $memory    = -1;
        public $loadAverage = -1;
        public $swap = -1;
        public $buffers = -1;
        public $cached = -1;

        public function serializeToRaw(){

            return $this->timestamp   . ' '
                .  $this->cpuUser     . ' '
                .  $this->cpuSystem   . ' '
                .  $this->memory      . ' '
                .  $this->loadAverage . ' '
                .  $this->swap        . ' '
                .  $this->buffers     . ' '
                .  $this->cached;
        }

        public function unSerializeFromRaw($raw){

            $rows = explode(' ', trim($raw));
            isset($rows[0]) && ($this->timestamp = intval($rows[0]));
            isset($rows[1]) && ($this->cpuUser = floatval($rows[1]));
            isset($rows[2]) && ($this->cpuSystem = floatval($rows[2]));
            isset($rows[3]) && ($this->memory = intval($rows[3]));
            isset($rows[4]) && ($this->loadAverage = floatval($rows[4]));
            isset($rows[5]) && ($this->swap = intval($rows[5]));
            isset($rows[6]) && ($this->buffers = intval($rows[6]));
            isset($rows[7]) && ($this->cached = intval($rows[7]));
        }

        // command of "top -d 0.01 -n 2", cpu of "-n 1" is not correct
        public function unSerializeFromTopOutput($raw){

            if($raw){
                // ubuntu will lost data, so use the first raw, while centos is ok
                $regex = '/.*?load average\:.*?([0-9\.]+), ([0-9\.]+), ([0-9\.]+).*?Cpu\(s\)\:.*? ([0-9\.]+)\%?.*? ([0-9\.]+)\%?.*? ([0-9\.]+)\%?.*? ([0-9\.]+)\%?.*? ([0-9\.]+)\%?.*? ([0-9\.]+)\%?.*? ([0-9\.]+)\%?.*? ([0-9\.]+)\%?.*? ([0-9]+)k?.*? ([0-9]+)k?.*? ([0-9]+)k?.*? ([0-9]+)k?.*? ([0-9]+)k?.*? ([0-9]+)k?.*? ([0-9]+)k?.*? ([0-9]+)k?/s';
                // sometimes $m has data but preg_match_all return false
                preg_match_all($regex,$raw,$m);
                if(isset($m[19][0])){
                    $this->timestamp = time();
                    $this->cpuUser = floatval(isset($m[4][1]) ? $m[4][1] : $m[4][0]);
                    $this->cpuSystem = floatval(isset($m[5][1]) ? $m[5][1] : $m[5][0]);
                    $this->loadAverage = floatval(isset($m[1][1]) ? $m[1][1] : $m[1][0]);
                    $this->memory = intval(isset($m[13][1]) ? $m[13][1] : $m[13][0]);
                    $this->swap = intval(isset($m[17][1]) ? $m[17][1] : $m[17][0]);
                    $this->buffers = intval(isset($m[15][1]) ? $m[15][1] : $m[15][0]);
                    $this->cached = intval(isset($m[19][1]) ? $m[19][1] : $m[19][0]);
                    return true;
                }
            }

            return false;
        }
    }

    abstract class DataManager {

        public static $saveDir = null ;
        public static $timezone = 28800;

        public static function save($id, DataBase $data){

            $dir = rtrim(self::$saveDir ? : __DIR__ . '/db','/\\') . '/';
            if(!file_exists($dir)){
                mkdir($dir);
            }
            $file =  $dir . $id . '_' . date('Ymd',$data->timestamp > 0 ? $data->timestamp : time()) . '.tbl';
            $content = $data->serializeToRaw();
            file_put_contents($file, $content. "\n", FILE_APPEND);
        }

        public static function get($id, $modelName, $startTime, $endTime = null){

            $dir = rtrim(self::$saveDir ? : __DIR__ . '/db','/\\');
            $endTime = $endTime ?: time();
            $list = [];
            $s = $startTime;
            //  make $e day end of $endTime
            $e = $endTime - (($endTime + self::$timezone ) % 86400) + 86400;
            while( $s < $e){
                $s += 86400;
                $file = $dir . '/' . $id . '_' . date('Ymd',$startTime) . '.tbl';
                if(!file_exists($file)){
                    continue;
                }
                if($fp = fopen($file,'r')){
                    while($row = fgets($fp)){
                        /* @var DataBase $model */
                        $model = new $modelName();
                        $model->unSerializeFromRaw($row);
                        if($model && $model->timestamp>= $startTime && $model->timestamp <= $endTime){
                            $list[] = $model;
                        }
                    }
                    fclose($fp);
                }
            }

            return $list;
        }
    }

    abstract class Curl {

        public static function request($url, $params = null, $isPost = true, $files = null) {

            $opts            = [ CURLOPT_RETURNTRANSFER => true,
                                 CURLOPT_HEADER         => false,
                                 CURLOPT_SSL_VERIFYPEER => false,
                                 CURLOPT_SSL_VERIFYHOST => false ];
            if( $isPost ){
                $opts[CURLOPT_POST]       = true;
                if($files){
                    $opts = self::customPostFields($params, $files) + $opts;
                }else if($params){
                    $opts[CURLOPT_POSTFIELDS] = http_build_query($params);
                }
            }else{
                if($params){
                    $url .= '?' . http_build_query($params);
                }
            }
            $ch = curl_init($url);
            curl_setopt_array($ch, $opts);
            $ret = curl_exec($ch);
            $info = curl_getinfo($ch);
            curl_close($ch);
            if( ( false !== $ret ) && ( 200 == $info['http_code'] ) ){
                return $ret;
            }else{
                return false;
            }
        }

        protected static function customPostFields(array $assoc = [], array $files = []) {

            static $disallow = ["\0", "\"", "\r", "\n"];
            $body = null;
            foreach ($assoc as $k => $v) {
                $k      = str_replace($disallow, "_", $k);
                $body[] = implode("\r\n", [ "Content-Disposition: form-data; name=\"{$k}\"", "", filter_var($v), ]);
            }
            foreach ($files as $k => $v) {
                switch (true) {
                    case false === $v = realpath(filter_var($v)):
                    case !is_file($v):
                    case !is_readable($v):
                        continue;
                }
                $data = file_get_contents($v);
                $v = call_user_func("end", explode(DIRECTORY_SEPARATOR, $v));
                $k = str_replace($disallow, "_", $k);
                $v = str_replace($disallow, "_", $v);
                $body[] = implode("\r\n", [ "Content-Disposition: form-data; name=\"{$k}\"; filename=\"{$v}\"", "Content-Type: application/octet-stream", "", $data, ]);
            }
            do {
                $boundary = "---------------------" . md5(mt_rand() . microtime());
            } while (preg_grep("/{$boundary}/", $body));
            array_walk($body, function (&$part) use ($boundary) {
                $part = "--{$boundary}\r\n{$part}";
            });
            $body[] = "--{$boundary}--";
            $body[] = "";
            return [CURLOPT_POST       => true,
                    CURLOPT_POSTFIELDS => implode("\r\n", $body),
                    CURLOPT_HTTPHEADER => [
                        "Expect: 100-continue",
                        "Content-Type: multipart/form-data; boundary={$boundary}",
                    ]];
        }
    }
}