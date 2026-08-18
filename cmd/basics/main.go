package main

import (
	"Ti28282/Redis-Study/internal/redis"
	"fmt"
	"time"
)

func main() {
	rdb := redis.NewClient()
	ctx := redis.GetContext()

	// STRINGS
	fmt.Println("=== STRING ===")
	rdb.Set(ctx, "name", "Alice", 0)
	rdb.Set(ctx, "counter", "10", 0)
	rdb.Incr(ctx, "counter") // result counter is 11

	rdb.Set(ctx, "temp", "value", 5*time.Second) // TTL 5 sec

	name, _ := rdb.Get(ctx, "name").Result()
	counter, _ := rdb.Get(ctx, "counter").Int()

	fmt.Println(name, counter)
	// > === HASH === <
	fmt.Println("\n=== HASH ===")
	rdb.HSet(ctx, "user:1001", []string{"user", "boob", "email", "bob@gmail.com", "age", "25"})
	rdb.HIncrBy(ctx, "user:1001", "age", 25) // age is 50

	AllUsers, _ := rdb.HGetAll(ctx, "user:1001").Result()

	fmt.Println(AllUsers)

	email, _ := rdb.HGet(ctx, "user:1001", "email").Result()

	fmt.Println(email)

	// > === LIST === <

	rdb.RPush(ctx, "tasks", "task1", "task2", "task3")
	rdb.LPush(ctx, "tasks", "urgent_task")

	tasks, _ := rdb.LRange(ctx, "tasks", 0, -1).Result()
	fmt.Printf("Tasks: %v\n", tasks)

	task, _ := rdb.LPop(ctx, "tasks").Result()
	fmt.Printf("Popped task: %s\n", task)

}
