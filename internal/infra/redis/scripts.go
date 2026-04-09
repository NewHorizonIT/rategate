package redis

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

type Script struct {
	script *redis.Script
}

func NewScript(lua string) *Script {
	return &Script{
		script: redis.NewScript(lua),
	}
}

func (s *Script) Run(ctx context.Context, client redis.Cmdable, keys []string, args ...interface{}) (interface{}, error) {
	result, err := s.script.Run(ctx, client, keys, args...).Result()
	return result, err
}

func MustLoadScript(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}
