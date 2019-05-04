package common

import (
	"github.com/go-redis/redis"
)

var DefaultRedis *redis.Client

func InitRedis(addr, passwd string) error {
	opt := &redis.Options{
		Addr:     addr,
		Password: passwd,
	}

	DefaultRedis = redis.NewClient(opt)
	_, err := DefaultRedis.Ping().Result()
	return err
}
