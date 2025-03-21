package dal

// 初始化连接redis

import (
	"context"

	"github.com/MakiJOJO/douyin-mall-echo/app/cart/config"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func RedisInit() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.GlobalConfig.Redis.Address,
		Username: config.GlobalConfig.Redis.Username,
		Password: config.GlobalConfig.Redis.Password,
		DB:       config.GlobalConfig.Redis.DB,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
