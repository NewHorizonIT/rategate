package ratelimit

import "errors"

// Define custom errors for the rate limit package
var (
	ErrEngineNotFound = errors.New("rate limit engine not found")
)
