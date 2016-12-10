# 使用Docker构建开发环境

前言：为了熟悉并跟进Docker，我们使用Docker来构建我们的开发环境。


## MySQL

第一步我们构建了mysql的运行环境，
第二步我们构建了mysql的主从复制。

这里我们用到了几个命令：

- 构建一个Docker镜像

```
docker build -t yaoguais/devel:1.0.0 .
```

- 使用docker-compose构建并启动镜像

```
sudo docker-compose build
sudo docker-compose up
```

- 查看并登录容器

```
docker ps
docker exec -it 085cfe0d51c0 bash
```

- 登录MySQL控制台

```
mysql --protocol=tcp -uroot -pyaoguais_2014 -P3306
```

- 配置主从复制

```
在宿主机中执行下面两条语句
echo "GRANT REPLICATION SLAVE ON *.* to 'slaves'@'%' identified by 'yaoguais_2014_2'; \
     flush privileges;"  \
	| mysql --protocol tcp -h127.0.0.1 -P3306 -uroot -pyaoguais_2014

echo "change master to master_host='127.0.0.1', master_port=3306, \
      master_user='slaves',master_password='yaoguais_2014_2', \
      master_log_file='binlog.000001',master_log_pos=0; \
      STOP SLAVE IO_THREAD FOR CHANNEL '';
      stop slave; \
      select sleep(1); \
      start slave;" \
    | mysql --protocol tcp -h127.0.0.1 -P3307 -uroot -pyaoguais_2014
```
