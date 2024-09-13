package redis_utils

import (
	"dominant/infrastructure/config"
	"github.com/go-redis/redis/v8"
	"time"
)

//
// @Author yfy2001
// @Date 2024/8/15 22 15
//

var GlobalRedisClient *redis.Client

func NewRedisClient(config config.Configuration) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         config.Redis.URL,
		Password:     config.Redis.Password,
		DB:           config.Redis.DB,
		DialTimeout:  time.Duration(config.Redis.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.Redis.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Redis.WriteTimeout) * time.Second,
		PoolSize:     config.Redis.PoolSize,
		PoolTimeout:  time.Duration(config.Redis.PoolTimeout) * time.Second,
	})
	return redisClient
}

func init() {
	GlobalRedisClient = NewRedisClient(config.Config)
}
