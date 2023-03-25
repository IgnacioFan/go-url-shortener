package redis

import (
	"fmt"
	"go-url-shortener/deployment/config"

	"github.com/go-redis/redis"
)

func InitClient() *redis.Client {
	config, err := config.New()
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("redis error:", err)
	}
	fmt.Println(pong)

	return client
}
