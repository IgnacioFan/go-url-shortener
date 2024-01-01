package redis

import (
	"errors"
	"fmt"
	"go-url-shortener/internal/util"
	"time"

	"github.com/go-redis/redis"
)

type Cache interface {
	Get(key string) (string, error) 
	Set(key, val string) error
}

type Impl struct {
	Client *redis.Client
}

func InitCache() (Cache, error) {
	redis := redis.NewClient(&redis.Options{Addr: util.RedisAddr()})
	pong, err := redis.Ping().Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(pong)
	return &Impl{Client: redis}, nil
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
