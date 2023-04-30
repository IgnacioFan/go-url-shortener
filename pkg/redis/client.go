package redis

type RedisClient interface {
	Get(key string) (string, error)
	Set(key, val string) error
	Del(key string) error
}
