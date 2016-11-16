#!/bin/bash
# yaoguais <newtopstdio@163.com>
# init the files structure

### mysql master
mkdir -p ./mysql/var/log/mysql
mkdir -p ./mysql/var/lib/mysql
cp -a ../mysql-master/etc ./mysql/

### mysql slave
mkdir -p ./mysql-slave/var/log/mysql
mkdir -p ./mysql-slave/var/lib/mysql
cp -a ../mysql-slave/etc ./mysql-slave/

## redis master
mkdir -p ./redis/var/lib/redis
mkdir -p ./redis/var/log/redis
cp -a ../redis/etc ./redis/

## redis slave
mkdir -p ./redis-slave/data
mkdir -p ./redis-slave/var/log/redis
cp -a ../redis-slave/etc ./redis/

## nginx
mkdir -p ./nginx/var/log/nginx
cp -a ../nginx/etc ./nginx/

## php-fpm
mkdir -p ./php-fpm/var/log/php
cp -a ../php-fpm/usr ./php-fpm/usr/

