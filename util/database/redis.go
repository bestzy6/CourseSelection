package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var RedisClient *redis.Client

// InitRedis 初始化Redis链接
func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:       addrRedis,
		Password:   pwd,
		DB:         dbnum,
		MaxRetries: 1,
	})
	timeout, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		log.Fatalln("link redis failed:", err)
	}
	RedisClient = client
}
