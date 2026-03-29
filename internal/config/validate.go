package config

import "errors"

func validateConfig(cfg *Config) error {
	if cfg.Server.Port == 0 {
		return errors.New("server.port is required")
	}

	if cfg.Redis.Host == "" {
		return errors.New("redis.host is required")
	}

	if cfg.RateLimit.Requests <= 0 {
		return errors.New("rateLimit.requests must > 0")
	}

	return nil
}
