#!/bin/bash
# history of build the dockers


docker build -t yaoguais/devel:1.0.0 .
mysql --protocol=tcp -uroot -pyaoguais_2014 -P3306

# grant mysql slave users
echo "GRANT REPLICATION SLAVE ON *.* to 'slaves'@'%' identified by 'yaoguais_2014_2'; \
     flush privileges;"  \
	| mysql --protocol tcp -h127.0.0.1 -P3306 -uroot -pyaoguais_2014

# be a slave of mysql master
# master_host be "mysql" cause by the "--links" param
# this command should run the first time only
echo "/*STOP SLAVE IO_THREAD FOR CHANNEL ''; reset slave;*/ \
      change master to master_host='127.0.0.1', master_port=3306, \
      master_user='slaves',master_password='yaoguais_2014_2', \
      master_log_file='binlog.000001',master_log_pos=0; \
      STOP SLAVE IO_THREAD FOR CHANNEL '';
      stop slave; \
      select sleep(1); \
      start slave;" \
    | mysql --protocol tcp -h127.0.0.1 -P3307 -uroot -pyaoguais_2014

apt-get install -y git wget strace vim net-tools tcpdump zip unzip tree
