<?php
/**
 * Created by PhpStorm.
 * User: yaoguai
 * Date: 15-7-5
 * Time: 下午3:30
 */

namespace fatty;

class MsgPack implements IPack{

    public function pack(&$data){
        return msgpack_pack($data);
    }

    public function unpack(&$data){
        return msgpack_unpack($data);
    }
}