package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/NewHorizonIT/rategate/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewClient(cnf config.RedisConfig) *redis.Client {
	Address := fmt.Sprintf("%v:%v", cnf.Host, cnf.Port)
	return redis.NewClient(&redis.Options{
		Addr:         Address,
		Password:     cnf.Password,
		DB:           cnf.DB,
		MinIdleConns: 10,
		DialTimeout:  2 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	})
}

func Ping(ctx context.Context, client *redis.Client) error {
	return client.Ping(ctx).Err()
}
