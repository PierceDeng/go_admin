package config

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var RedisTemplate *redis.Client

func InitRedis() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"), // Redis 服务器地址，格式为 "host:port"
		Password: viper.GetString("redis.password"),                                   // 如果 Redis 设置了密码，在这里填写
		DB:       viper.GetInt("redis.db"),                                            // 使用的数据库编号，默认为 0
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
