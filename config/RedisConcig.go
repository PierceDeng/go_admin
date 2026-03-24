package config

import (
	"github.com/redis/go-redis/v9"
)

var RedisTemplate *redis.Client

func InitRedis() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.30.128:6379", // Redis 服务器地址，格式为 "host:port"
		Password: "",                    // 如果 Redis 设置了密码，在这里填写
		DB:       0,                     // 使用的数据库编号，默认为 0
	})

	RedisTemplate = rdb
}

func CloseRedis() {

	if RedisTemplate != nil {
		err := RedisTemplate.Close()
		if err != nil {
			return
		}
	}

}
