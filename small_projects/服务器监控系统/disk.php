<?php

namespace watcher {

    class DiskProtocol extends Protocol {

        const COMMAND_DISK = 200;

        public function handle(User $user, Packet $packet){

        }
    }

    class DiskTick extends TickBase {

        protected $sshConfig;

        public function __construct($config){

            parent::__construct($config);
            SshManager::getInstance()->addConnection($config);
            $this->sshConfig = $config;
        }

        public function tick($id, $server){

            /* @var Ssh $ssh */
            $ssh = SshManager::getInstance()->getConnection($this->sshConfig);
            if($ret = $ssh->exec('df')){
                $data = new DiskData();
                if($data->unSerializeFromDfOutput($ret)){
                    DataManager::save($this->getId(), $data);
                    $userManager =  UserManager::getInstance();
                    $userManager->each(function($user) use($userManager, $data){
                        if($user->auth){
                            $packet = new Packet();
                            $packet->command = DiskProtocol::COMMAND_DISK;
                            $packet->data = ['id' => $this->getId(), 'info' => $data->serializeToRaw()];
                            $userManager->sendMessage($user, $packet->serialize());
                        }
                    });
                }else{
                    Logger::log('parse df output error, ticker ' . $this->getId(), Logger::LEVEL_ERROR);
                }
            }else{
                Logger::log('ssh exec error, '. $ssh->getLastError(), Logger::LEVEL_ERROR);
                $this->setStop(true);
            }
        }
    }

    class DiskData extends DataBase {

        public $diskInfo; // 用逗号隔开每个磁盘的信息

        public function serializeToRaw(){

            return date('Y-m-d H:i:s', $this->timestamp) . ' ' . $this->diskInfo;
        }

        public function unSerializeFromRaw($raw){

            $rows = explode(' ', trim($raw), 2);
            isset($rows[0]) && ($this->timestamp = intval($rows[0]));
            isset($rows[1]) && ($this->diskInfo = floatval($rows[1]));
        }

        public function unSerializeFromDfOutput($raw){

            if(!$raw){
                return false;
            }
            $rows = explode("\n",$raw);
            $row = '';
            foreach($rows as $r){
                $row .=  trim($r) . ';';
            }
            $this->timestamp = time();
            $this->diskInfo = $row;
            return true;
        }
    }
}