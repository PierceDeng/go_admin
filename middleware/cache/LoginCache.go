package cache

import (
	"context"
	"go_admin/config"
	"go_admin/middleware/exception"
)

func SetSysToken(token string, userId uint64) {
	tokenKey := "system:user:login:" + token
	config.RedisTemplate.Set(context.Background(), tokenKey, userId, 60)
}

func GetSysToken(token string) uint64 {
	tokenKey := "system:user:login:" + token
	userId, err := config.RedisTemplate.Get(context.Background(), tokenKey).Uint64()
	if err != nil {
		return userId
	}
	panic(exception.NewBizException(10002, "Token已失效"))
}
