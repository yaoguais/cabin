<?php

interface IQueue{
	/**
	 * 入队
	 * @param string $message
	 * @return bool
	 */
	public function push($message);
	/**
	 * 出队
	 * @param void
	 * @return string $message | false
	 */
	public function pop();
	/**
	 * 队首元素
	 * @param void
	 * @return string $message | false
	 */
	public function front();
	/**
	 * 队尾元素
	 * @param void
	 * @return string $message | false
	 */
	public function back();
	/**
	 * 是否空队
	 * @param void
	 * @return bool
	 */
	public function isEmpty();
	/**
	 * 队列长度
	 * @param void
	 * @return int
	 */
	public function size();
}