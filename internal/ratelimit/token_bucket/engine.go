package tokenbucket

import (
	"context"
	"fmt"

	"github.com/NewHorizonIT/rategate/internal/infra/redis"
	"github.com/NewHorizonIT/rategate/internal/ratelimit"
	"github.com/NewHorizonIT/rategate/pkg/helper"
)

type TokenBucketEngine struct {
	redis redis.RedisRepo
}

func New(redis redis.RedisRepo) *TokenBucketEngine {
	return &TokenBucketEngine{redis: redis}
}

func (t *TokenBucketEngine) Allow(
	ctx context.Context,
	policy ratelimit.Policy,
	req ratelimit.Request,
) (bool, int64, error) {

	key := helper.BuildKey(req.Tenant, req.User, req.Endpoint)

	res, err := t.redis.Script.Run(ctx, t.redis.Client, []string{key}, policy.Limit, policy.Window)
	if err != nil {
		return false, 0, fmt.Errorf("lua exec failed: %w", err)
	}

	values, ok := res.([]interface{})
	if !ok || len(values) != 2 {
		return false, 0, fmt.Errorf("invalid lua response")
	}

	allowed := values[0].(int64)
	count := values[1].(int64)

	return allowed == 1, count, nil
}
