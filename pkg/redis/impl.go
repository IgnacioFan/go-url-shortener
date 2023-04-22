package redis

import (
	"errors"
	"fmt"
	"go-url-shortener/config"
	"time"

	"github.com/go-redis/redis"
)

type Impl struct {
	Client *redis.Client
}

func InitClient(config *config.Config) RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
	})
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("redis error:", err)
	}
	fmt.Println(pong)

	return &Impl{
		Client: client,
	}
}

func (i *Impl) Get(key string) (string, error) {
	if str, err := i.Client.Get(key).Result(); err == redis.Nil {
		return "", errors.New("No entry")
	} else {
		return str, err
	}
}

func (i *Impl) Set(key, val string) error {
	expiration := time.Duration(24 * time.Hour)
	if _, err := i.Client.Set(key, val, expiration).Result(); err != nil {
		return err
	}
	return nil
}