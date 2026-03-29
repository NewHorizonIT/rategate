package config

import "log"

func SetupConfig() *Config {
	loader := NewViperProvider()
	cfg, err := loader.Load()
	if err != nil {
		log.Fatal()
	}
	return cfg
}
