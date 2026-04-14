package ristretto

import (
	"time"

	"github.com/NewHorizonIT/rategate/internal/config"
	"github.com/dgraph-io/ristretto/v2"
)

// Struct L1 Cache
type L1Cache struct {
	cache *ristretto.Cache[string, string]
	ttl   time.Duration
}

// Register L1 cache
// Using Ristretto as the in-memory cache implementation
func RegisterL1Cache(conf config.CacheConfig) (IInMemCache, error) {
	cnfRistretto := ristretto.Config[string, string]{
		NumCounters: int64(conf.NumCounters),
		MaxCost:     int64(conf.MaxCost),
		BufferItems: int64(conf.BufferItems),
	}
	cache, err := ristretto.NewCache(&cnfRistretto)
	if err != nil {
		return nil, err
	}
	return &L1Cache{cache: cache}, nil
}

// Delete implements IInMemCache.
func (l *L1Cache) Delete(key string) error {
	panic("unimplemented")
}

// Get implements IInMemCache.
func (l *L1Cache) Get(key string) (string, error) {
	panic("unimplemented")
}

// Set implements IInMemCache.
func (l *L1Cache) Set(key string, value string) error {
	panic("unimplemented")
}
