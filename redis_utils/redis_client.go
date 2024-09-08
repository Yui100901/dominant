package redis_utils

import (
	"dominant/config"
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
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     20,
		PoolTimeout:  30 * time.Second,
	})
	return redisClient
}

func init() {
	GlobalRedisClient = NewRedisClient(config.Config)
}
