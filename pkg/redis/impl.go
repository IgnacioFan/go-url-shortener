package redis

import (
	"errors"
	"fmt"
	"go-url-shortener/deployment/env"
	"time"

	"github.com/go-redis/redis"
)

type Impl struct {
	Client *redis.Client
}

func InitClient() (RedisClient, error) {
	client := redis.NewClient(&redis.Options{Addr: env.RedisAddr()})
	pong, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(pong)
	return &Impl{Client: client}, nil
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

func (i *Impl) Del(key string) error {
	if _, err := i.Client.Del(key).Result(); err != nil {
		return err
	}
	return nil
}
