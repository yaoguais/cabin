<?php
/**
 * Created by PhpStorm.
 * User: yaoguai
 * Date: 15-7-5
 * Time: 下午3:49
 */

namespace fatty;

class PhpPack implements IPack{

    public function pack(&$data){
        return serialize($data);
    }

    public function unpack(&$data){
        return unserialize($data);
    }
}