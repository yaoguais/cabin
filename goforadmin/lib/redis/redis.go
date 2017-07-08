package redis

import (
	"gopkg.in/redis.v5"
)

var Redis *redis.Client

type RedisConf struct {
	Address  string
	Password string
	Database int
}

func Connect(conf *RedisConf) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		Password: conf.Password,
		DB:       conf.Database,
	})

	if err := client.Ping().Err(); err != nil {
		panic("connect to redis failed " + err.Error())
	}

	if Redis == nil {
		Redis = client
	}

	return client
}
