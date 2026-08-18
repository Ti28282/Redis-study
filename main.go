package main

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)
	if err := rdb.Set(ctx, "message:1", "hello world", 60*time.Second).Err(); err != nil {
		log.Fatalf("Error: %s", err)
	}
	rdb.Set(ctx, "message", "hello my name is Timur", 60*time.Second)

	rdb.TTL(ctx, "message")
}
