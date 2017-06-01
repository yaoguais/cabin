#!/bin/bash
# yaoguais <newtopstdio@163.com>
# init the files structure

mkdir -p ./var/lib/mysql
mkdir -p ./var/log/mysql
cd ./var/log/mysql
touch error.log slow.log general.log
cd -

mkdir ./etc/mysql/conf.d

