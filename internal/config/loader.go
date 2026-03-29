package config

import (
	"fmt"

	"github.com/NewHorizonIT/rategate/pkg"
	"github.com/spf13/viper"
)

type ViperProvider struct{}

// Initializer for ViperProvider
func NewViperProvider() *ViperProvider {
	return &ViperProvider{}
}

func (p *ViperProvider) Load() (*Config, error) {
	v := viper.New()

	bindFlags(v)
	bindEnv(v)

	env := pkg.GetEnv("APP_ENV", "dev")

	// load config file
	if err := loadFile(v, env); err != nil {
		return nil, fmt.Errorf("error loading config file: %w", err)
	}

	// Unmarshal config into struct
	cfg := Config{}
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	// Override with env vars and flags
	applyEnvOverrides(v, &cfg)

	// Dynamic config
	v.WatchConfig()

	// Validate config
	err := validateConfig(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error validating config: %w", err)
	}
	return &cfg, nil
}
