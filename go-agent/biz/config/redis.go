package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)


func InitRedisClient(ctx context.Context) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
