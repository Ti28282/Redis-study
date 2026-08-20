package redisClient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func GetContext() context.Context {
	return ctx
}
