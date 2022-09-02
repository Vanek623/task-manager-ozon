package config

import "time"

// Redis конфиги редиса
type Redis struct {
	Host         string
	DB           int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var redis Redis

func GetRedisConfig() Redis {
	return redis
}

func init() {
	redis.Host = "localhost:6379"
	redis.DB = 0
	redis.ReadTimeout = 100 * time.Millisecond
	redis.WriteTimeout = 200 * time.Millisecond
}
