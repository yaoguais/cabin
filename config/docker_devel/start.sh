#!/bin/bash
# yaoguais <newtopstdio@163.com>
# start develop cluster

sudo docker-compose up -d
sleep 2
echo "start slave;" | mysql --protocol tcp -h127.0.0.1 -P3307 -uroot -pyaoguais_2014
echo
