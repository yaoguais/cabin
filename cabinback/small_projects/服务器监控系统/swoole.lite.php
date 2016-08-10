<?php
class swoole_server {
	/* methods */
	public function __construct($serv_host, $serv_port, $serv_mode = NULL, $sock_type = NULL) {
	}
	public function set($zset) {
	}
	public function start() {
	}
	public function send($conn_fd, $send_data, $from_id = NULL) {
	}
	public function sendto($ip, $port, $send_data = NULL) {
	}
	public function sendwait($conn_fd, $send_data) {
	}
	public function exist($conn_fd) {
	}
	public function sendfile($conn_fd, $filename) {
	}
	public function close($fd) {
	}
	public function task($data, $worker_id) {
	}
	public function taskwait($data, $timeout = NULL, $worker_id = NULL) {
	}
	public function finish($data) {
	}
	public function addlistener($host, $port, $sock_type) {
	}
	public function listen($host, $port, $sock_type) {
	}
	public function reload() {
	}
	public function shutdown() {
	}
	public function hbcheck($from_id) {
	}
	public function heartbeat($from_id) {
	}
	public function handler($ha_name, $cb) {
	}
	public function on($ha_name, $cb) {
	}
	public function connection_info($fd, $from_id) {
	}
	public function connection_list($start_fd, $find_count) {
	}
	public function getClientInfo($fd, $from_id) {
	}
	public function getClientList($start_fd, $find_count) {
	}
	public function addtimer($interval) {
	}
	public function deltimer($interval) {
	}
	public function gettimer() {
	}
	public function after() {
	}
	public function tick() {
	}
	public function clearTimer() {
	}
	public function sendmessage() {
	}
	public function addprocess() {
	}
	public function stats() {
	}
	public function bind($fd, $uid) {
	}
}
class swoole_connection_iterator implements Iterator, Traversable, Countable {
	/* methods */
	public function rewind() {
	}
	public function next() {
	}
	public function current() {
	}
	public function key() {
	}
	public function valid() {
	}
	public function count() {
	}
}
class swoole_client {
	/* properties */
	public $errCode = "0";
	public $sock = "0";

	/* methods */
	public function __construct() {
	}
	public function __destruct() {
	}
	public function set() {
	}
	public function connect() {
	}
	public function recv() {
	}
	public function send() {
	}
	public function sendfile() {
	}
	public function sendto() {
	}
	public function isConnected() {
	}
	public function getsockname() {
	}
	public function getpeername() {
	}
	public function close() {
	}
	public function on() {
	}
}
class swoole_process {
	/* methods */
	public function __construct() {
	}
	public function __destruct() {
	}
	public static function wait() {
	}
	public static function signal() {
	}
	public static function kill() {
	}
	public static function daemon() {
	}
	public static function setaffinity() {
	}
	public function useQueue() {
	}
	public function start() {
	}
	public function write() {
	}
	public function close() {
	}
	public function read() {
	}
	public function push() {
	}
	public function pop() {
	}
	/*public function exit() {
	}*/
	public function exec() {
	}
	public function name() {
	}
}
class swoole_table implements Iterator, Traversable, Countable {
	/* constants */
	const TYPE_INT = "1";
	const TYPE_STRING = "7";
	const TYPE_FLOAT = "6";

	/* methods */
	public function __construct($table_size) {
	}
	public function column($name, $type = NULL, $size = NULL) {
	}
	public function create() {
	}
	public function set($key, $value) {
	}
	public function get($key) {
	}
	public function count() {
	}
	public function del($key) {
	}
	public function exist($key) {
	}
	public function incr($key, $column, $incrby = NULL) {
	}
	public function decr($key, $column, $decrby = NULL) {
	}
	public function lock() {
	}
	public function unlock() {
	}
	public function rewind() {
	}
	public function next() {
	}
	public function current() {
	}
	public function key() {
	}
	public function valid() {
	}
}
class swoole_lock {
	/* methods */
	public function __construct() {
	}
	public function __destruct() {
	}
	public function lock() {
	}
	public function trylock() {
	}
	public function lock_read() {
	}
	public function trylock_read() {
	}
	public function unlock() {
	}
}
class swoole_atomic {
	/* methods */
	public function __construct() {
	}
	public function add() {
	}
	public function sub() {
	}
	public function get() {
	}
	public function set() {
	}
	public function cmpset() {
	}
}
class swoole_http_server extends swoole_server {
	/* properties */
	private $global = "0";

	/* methods */
	public function on($ha_name, $cb) {
	}
	public function setglobal() {
	}
	public function start() {
	}
	public function __construct($serv_host, $serv_port, $serv_mode = NULL, $sock_type = NULL) {
	}
	public function set($zset) {
	}
	public function send($conn_fd, $send_data, $from_id = NULL) {
	}
	public function sendto($ip, $port, $send_data = NULL) {
	}
	public function sendwait($conn_fd, $send_data) {
	}
	public function exist($conn_fd) {
	}
	public function sendfile($conn_fd, $filename) {
	}
	public function close($fd) {
	}
	public function task($data, $worker_id) {
	}
	public function taskwait($data, $timeout = NULL, $worker_id = NULL) {
	}
	public function finish($data) {
	}
	public function addlistener($host, $port, $sock_type) {
	}
	public function listen($host, $port, $sock_type) {
	}
	public function reload() {
	}
	public function shutdown() {
	}
	public function hbcheck($from_id) {
	}
	public function heartbeat($from_id) {
	}
	public function handler($ha_name, $cb) {
	}
	public function connection_info($fd, $from_id) {
	}
	public function connection_list($start_fd, $find_count) {
	}
	public function getClientInfo($fd, $from_id) {
	}
	public function getClientList($start_fd, $find_count) {
	}
	public function addtimer($interval) {
	}
	public function deltimer($interval) {
	}
	public function gettimer() {
	}
	public function after() {
	}
	public function tick() {
	}
	public function clearTimer() {
	}
	public function sendmessage() {
	}
	public function addprocess() {
	}
	public function stats() {
	}
	public function bind($fd, $uid) {
	}
}
class swoole_http_response {
	/* methods */
	public function cookie() {
	}
	public function rawcookie() {
	}
	public function status() {
	}
	public function gzip() {
	}
	public function header() {
	}
	public function write() {
	}
	public function end() {
	}
}
class swoole_http_request {
	/* methods */
	public function rawcontent() {
	}
}
class swoole_buffer {
	/* methods */
	public function __construct() {
	}
	public function __destruct() {
	}
	public function substr() {
	}
	public function write() {
	}
	public function read() {
	}
	public function append() {
	}
	public function expand() {
	}
	public function clear() {
	}
}
class swoole_websocket_server extends swoole_http_server {
	/* methods */
	public function on($event_name, $callback) {
	}
	public function push($fd, $data, $opcode = NULL, $finish = NULL) {
	}
	public function setglobal() {
	}
	public function start() {
	}
	public function __construct($serv_host, $serv_port, $serv_mode = NULL, $sock_type = NULL) {
	}
	public function set($zset) {
	}
	public function send($conn_fd, $send_data, $from_id = NULL) {
	}
	public function sendto($ip, $port, $send_data = NULL) {
	}
	public function sendwait($conn_fd, $send_data) {
	}
	public function exist($conn_fd) {
	}
	public function sendfile($conn_fd, $filename) {
	}
	public function close($fd) {
	}
	public function task($data, $worker_id) {
	}
	public function taskwait($data, $timeout = NULL, $worker_id = NULL) {
	}
	public function finish($data) {
	}
	public function addlistener($host, $port, $sock_type) {
	}
	public function listen($host, $port, $sock_type) {
	}
	public function reload() {
	}
	public function shutdown() {
	}
	public function hbcheck($from_id) {
	}
	public function heartbeat($from_id) {
	}
	public function handler($ha_name, $cb) {
	}
	public function connection_info($fd, $from_id) {
	}
	public function connection_list($start_fd, $find_count) {
	}
	public function getClientInfo($fd, $from_id) {
	}
	public function getClientList($start_fd, $find_count) {
	}
	public function addtimer($interval) {
	}
	public function deltimer($interval) {
	}
	public function gettimer() {
	}
	public function after() {
	}
	public function tick() {
	}
	public function clearTimer() {
	}
	public function sendmessage() {
	}
	public function addprocess() {
	}
	public function stats() {
	}
	public function bind($fd, $uid) {
	}
}
class swoole_websocket_frame {
}
define('SWOOLE_BASE',4);
define('SWOOLE_THREAD',2);
define('SWOOLE_PROCESS',3);
define('SWOOLE_PACKET',16);
define('SWOOLE_IPC_UNSOCK',1);
define('SWOOLE_IPC_MSGQUEUE',2);
define('SWOOLE_IPC_CHANNEL',3);
define('SWOOLE_SOCK_TCP',1);
define('SWOOLE_SOCK_TCP6',3);
define('SWOOLE_SOCK_UDP',2);
define('SWOOLE_SOCK_UDP6',4);
define('SWOOLE_SOCK_UNIX_DGRAM',5);
define('SWOOLE_SOCK_UNIX_STREAM',6);
define('SWOOLE_TCP',1);
define('SWOOLE_TCP6',3);
define('SWOOLE_UDP',2);
define('SWOOLE_UDP6',4);
define('SWOOLE_UNIX_DGRAM',5);
define('SWOOLE_UNIX_STREAM',6);
define('SWOOLE_SOCK_SYNC',0);
define('SWOOLE_SOCK_ASYNC',1);
define('SWOOLE_SYNC',2048);
define('SWOOLE_ASYNC',1024);
define('SWOOLE_KEEP',4096);
define('SWOOLE_EVENT_READ',512);
define('SWOOLE_EVENT_WRITE',1024);
define('SWOOLE_VERSION','1.7.20');
define('SWOOLE_AIO_BASE',0);
define('SWOOLE_AIO_GCC',1);
define('SWOOLE_AIO_LINUX',2);
define('SWOOLE_FILELOCK',2);
define('SWOOLE_MUTEX',3);
define('SWOOLE_SEM',4);
define('SWOOLE_RWLOCK',1);
define('SWOOLE_SPINLOCK',5);
