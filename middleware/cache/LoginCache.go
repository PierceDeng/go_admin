package cache

import (
	"context"
	"go_admin/config"
	"time"
)

func SetSysToken(token string, userId uint64) {
	tokenKey := "system:user:login:" + token
	config.RedisTemplate.Set(context.Background(), tokenKey, userId, 60*time.Minute)
}

func GetSysToken(token string) uint64 {
	tokenKey := "system:user:login:" + token
	userId, _ := config.RedisTemplate.Get(context.Background(), tokenKey).Uint64()
	return userId
}
