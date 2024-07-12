package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var DB *redis.Client

func ConnectToRedis() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379", // Redis address
	})

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379", // Use the service name 'redis' defined in Docker Compose
	// 	Password: "",           // No password set
	// 	DB:       0,            // Use default DB
	// })

	// Test connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Could not connect to Redis:", err)
		return
	}
	DB = rdb
	fmt.Println("Connected to Redis:", pong)
}

func InsertIntoRedis(key, value string) error {
	err := DB.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		log.Println("Error in inserting to Redis")
		return err
	}
	return nil
}

func GetFromDB(key string) (string, error) {
	value, err := DB.Get(context.Background(), key).Result()
	if err != nil {
		log.Println("Error in getting the key from Redis")
		return value, err
	}
	return value, nil
}
