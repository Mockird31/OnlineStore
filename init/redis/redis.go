package redis

import (
	"fmt"
	"time"

	"github.com/Mockird31/OnlineStore/config"
	"github.com/gomodule/redigo/redis"
)

func NewRedisPool(cfg config.RedisConfig) *redis.Pool {
	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
