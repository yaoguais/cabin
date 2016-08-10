<?php

namespace watcher {

    class RequestProtocol extends TickProtocol {

        const COMMAND_GET = TopPacket::COMMAND_GET;

        public function handle(User $user, Packet $packet){

            $userManager = UserManager::getInstance();
            $tickManager = TickManager::getInstance();
            switch($packet->command){
                case self::COMMAND_GET:
                    if(true){
                        $list = [];
                        if(!$packet->data){
                            $idList = $tickManager->each(function($tick){
                                if($tick instanceof RequestTick){
                                    return $tick->getId();
                                }
                            });
                            if(is_array($idList)){
                                $start = $end = null;
                                foreach($idList as $id){
                                    if($id && ($data = $this->getList($id, 'watcher\\RequestData', $start, $end))){
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
            }
        }


        protected function parseStatDataToSimpleArray(DataBase $model){

            return [
                $model->timestamp,
                $model->time
            ];
        }
    }

    class RequestTick extends TickBase {

        protected $url;

        public function __construct($config){

            parent::__construct($config);
            $this->url = $config['url'];
        }

        public function tick($id, $server){

            $start = microtime(true);
            $ret = Curl::request($this->url);
            $time = microtime(true) - $start;
            if($ret && strpos($ret,'{"code":1') !==false){
                $data = new RequestData();
                $data->timestamp = time();
                $data->time = $time;
                DataManager::save($this->getId(),$data);
                (new RequestProtocol())->handleFlowStatData($this, $data);
            }
        }
    }

    class RequestData extends DataBase {

        public $time = -1;

        public function serializeToRaw(){

            return $this->timestamp   . ' '
            .  $this->time;
        }

        public function unSerializeFromRaw($raw){

            $rows = explode(' ', trim($raw));
            isset($rows[0]) && ($this->timestamp = intval($rows[0]));
            isset($rows[1]) && ($this->time = floatval($rows[1]));
        }
    }
}