package main

import (
	"Ti28282/Redis-Study/internal/redisClient"
	"fmt"

	"time"

	"github.com/redis/go-redis/v9"
)

func main() {

	rdb := redisClient.NewClient()
	ctx := redisClient.GetContext()

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

	// > === SET === <
	fmt.Println("> === SET === <")
	rdb.SAdd(ctx, "tags:post:1", "go", "redis", "backend")
	rdb.SAdd(ctx, "tags:post:2", "python", "redis", "api")

	common, _ := rdb.SInter(ctx, "tags:post:1", "tags:post:2").Result()
	fmt.Printf("Common tags: %v\n", common)

	isMember, _ := rdb.SIsMember(ctx, "tags:post:1", "go").Result()
	fmt.Printf("Has 'go' tag: %v\n", isMember)

	// > === SORTED SET === <
	fmt.Println("> === SORTED SET === <")

	rdb.ZAdd(ctx, "leaderboard", redis.Z{Score: 1500, Member: "player1"})
	rdb.ZAdd(ctx, "leaderboard", redis.Z{Score: 2300, Member: "player2"})
	rdb.ZAdd(ctx, "leaderboard", redis.Z{Score: 1800, Member: "player3"})

	top3, _ := rdb.ZRevRange(ctx, "leaderboard", 0, 2).Result()
	fmt.Printf("Top 3: %v\n", top3)

	rank, _ := rdb.ZRank(ctx, "leaderboard", "player1").Result()
	fmt.Printf("Player1 rank: %d\n", rank+1)

	// === TTL and delete ===
	fmt.Println("\n > === TTL & DELETE === <")
	ttl, _ := rdb.TTL(ctx, "temp").Result()
	fmt.Printf("Temp key TTL: %v\n", ttl)

	rdb.Del(ctx, "name")

	exists, _ := rdb.Exists(ctx, "name").Result()
	fmt.Printf("Name exists: %v\n", exists > 0)
}
