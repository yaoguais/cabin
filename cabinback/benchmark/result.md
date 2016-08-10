```
redis 40000 client "server 500 error"
mongo 50000 client "server 500 error"
            server "Resource temporarily unavailable"
mongo 40000 client "server 500 error"
                   "too many open files"
mongo 30000 ok Mean:143ms
redis 40000 client "server 500 error"
redis 30000 ok Mean:15ms
redis 40000 ok Mean:19ms
redis 40000 ok Mean:19ms
redis 50000 client "server 500 error"


finally-max:
echo:  90000
mongo: 30000
redis: 40000
mysql: 90000


finally-mean:
echo:  70ms 101ms
mongo: 143ms 146ms
redis: 19ms 19ms
mysql: 63ms 40ms

```

##  出现的问题及解决的办法  ##

```
1.cannot assign requested address 端口不够用，应该加速sock的释放或启动重用
2.too many open files 进程打开的文件描述符太多，加大能够打开的文件数量
3.apache max clients  Apache进程最大连接数太小
4.mongodb maxConn     MongoDB最大连接数太小
5.redis maxCliens     Redis最大连接数太小
6.mysql maxConnection Mysql最大连接数太小

```
