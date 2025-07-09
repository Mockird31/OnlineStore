package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/Mockird31/OnlineStore/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis(cfg *config.Config) (*redis.Client, error) {
	address := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:         address,
		Network:      "tcp",
		MinIdleConns: 5,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
