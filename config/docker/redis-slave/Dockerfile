FROM redis:3.0.7
MAINTAINER yaoguais <newtopstdio@163.com>

EXPOSE 6380
COPY ./etc/redis/redis.conf /etc/redis/redis.conf
CMD [ "redis-server", "/etc/redis/redis.conf" ]

