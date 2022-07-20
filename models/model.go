package models

import (
	"fmt"
	"whois-api/configer"

	"github.com/go-redis/redis"
)

var (
	RedisClient *redis.Client
)

func InitialRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", configer.Configer.Redis.Host, configer.Configer.Redis.Port),
		Password: configer.Configer.Redis.Password,
		DB:       0,
		PoolSize: 100,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("initial redis fail,%s", err))
	}
}
