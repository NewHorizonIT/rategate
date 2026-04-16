package ratelimit

import (
	"fmt"
)

var engines = map[string]Engine{}

func Register(name string, engine Engine) {
	engines[name] = engine
}

func Get(name string) (Engine, error) {
	if engine, exists := engines[name]; exists {
		return engine, nil
	}
	return nil, fmt.Errorf("%w: %s", ErrEngineNotFound, name)
}

// Usage:
// Register engines in main.go
// tokenBucketEngine := ratelimit.Register("TokenBucket", tokenbucket.New(redisClient))

// engine := Get("TokenBucket")
// result, err := engine.Allow(ctx, policy, req)
