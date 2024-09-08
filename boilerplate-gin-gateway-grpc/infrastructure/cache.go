package infrastructure

import (
	"context"
	"fmt"
	"os"

	"server/env"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func ConnRedis() {
	addr := fmt.Sprintf("%v:%v", env.NewEnv().REDIS_HOST, env.NewEnv().REDIS_PORT)
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: env.NewEnv().REDIS_USERNAME,
		Password: env.NewEnv().REDIS_PASSWORD,
		DB:       env.NewEnv().REDIS_DATABASE,
	})
	ctx := context.Background()
	if err := Redis.Ping(ctx).Err(); err != nil {
		fmt.Printf("redis: can't ping to redis - %v \n", err)
		os.Exit(1)
	}
	fmt.Println("redis: connection opened")
}
