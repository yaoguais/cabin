<?php
/**
 * Created by PhpStorm.
 * User: yaoguai
 * Date: 15-7-5
 * Time: 下午3:29
 */

namespace fatty;

interface IPack{
    /**
     * @param $data
     * @return string
     */
    public function pack(&$data);

    /**
     * @param string $data
     * @return mixed
     */
    public function unpack(&$data);
}