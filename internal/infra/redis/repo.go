package redis

import "github.com/redis/go-redis/v9"

type RedisRepo struct {
	Client *redis.Client
	Script *Script
}
