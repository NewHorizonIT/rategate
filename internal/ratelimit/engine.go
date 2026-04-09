package ratelimit

import "context"

type Engine interface {
	Allow(ctx context.Context, policy Policy, req Request) (Result, error)
}
