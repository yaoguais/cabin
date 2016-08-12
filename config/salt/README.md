## salt configuration

主机:

	主机名: salt1 IP地址: 192.168.1.241 作用:LB和salt-master机器
	主机名: salt2 IP地址: 192.168.1.242 作用:WEB服务器和salt-minion机器1 salt名称: web-server1
	主机名: salt3 IP地址: 192.168.1.243 作用:WEB服务器和salt-minion机器2 salt名称: web-server2

网路配置:

	三台机器网络配置只有IP地址不同，其它完全相同。
	首先配置网卡eth0,设置静态IP。
	然后配置网关/etc/sysconfig/network
	然后配置DNS服务器/etc/resolv.conf


配置salt-master机器的salt服务:

	添加yum源: /etc/yum.repos.d/saltstack.repo
	yum install salt-master
	
	vim /etc/salt/master
	修改以下字段
	interface: 192.168.1.249
	hash_type: sha256
	
	启动服务
	service salt-master start


配置salt-minion机器1的salt服务:

	添加yum源: /etc/yum.repos.d/saltstack.repo
	yum install salt-minion
	
	vim /etc/salt/minion
	修改以下字段
	id: web-server1
	master: 192.168.1.241
	hash_type: sha256

	启动服务
	service salt-minion start


我们使用同样的方式配置salt-minion机器2的salt服务。



通过上面的操作，我们已经配置好了3台机器的salt服务，然后我们开始测试3台服务器的连接。

	查看所有的minion机器:
	# salt-key -L
	Accepted Keys:
	Denied Keys:
	Unaccepted Keys:
	web-server1
	web-server2
	Rejected Keys:

	接受2台minion服务器:
	# salt-key -a web-server1
	# salt-key -a web-server2
	
	测试两台服务器的连接:
	# salt "web-server*" test.ping
	web-server1:
    	True
	web-server2:
    	True


然后我们配置2台minion服务器组件:

- 基础组件 base.sls
- NGINX nginx.sls
- PHP7 php7.sls
- Redis redis.sls
- MySql mysql.sls
- MongoDB mongodb.sls
- Postgresql postgresql.sls
- NodeJs nodejs.sls
- Supervisor supervisor.sls
- PhpMyAdmin phpmyadmin.sls
- Golang golang.sls
- Sync sync.sls
- Lvs lvs.sls

写完这些配置文件,我们依次安装这些组件.
所有的配置文件可以到[这里找到](https://github.com/Yaoguais/cabin/tree/master/config/salt/salt1_file_system).

    # salt 'web-server*' state.sls install.base,install.nginx,install.php7,install.redis,
    install.mysql,install.mongodb,install.postgresql,install.nodejs,install.supervisor,install.phpmyadmin,
    install.golang,install.sync
    # salt 'web-lb' state.sls install.lvs

运行完上面的这条命令,等待一段时间,所有的组件就安装完毕了.