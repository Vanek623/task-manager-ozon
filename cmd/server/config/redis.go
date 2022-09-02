package config

import (
	"time"

	redisPkg "github.com/go-redis/redis/v8"
)

var redis redisPkg.Options

func GetRedisConfig() redisPkg.Options {
	return redis
}

func init() {
	redis.Addr = "localhost:6379"
	redis.DB = 0
	redis.ReadTimeout = 100 * time.Millisecond
	redis.WriteTimeout = 200 * time.Millisecond
}
