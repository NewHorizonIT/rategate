package config

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func loadFile(v *viper.Viper, env string) error {
	v.SetConfigName(fmt.Sprintf("config.%s", env))
	v.SetConfigType("yaml")

	// Set path
	v.AddConfigPath("./configs")
	// Read config
	return v.ReadInConfig()
}

func applyEnvOverrides(v *viper.Viper, cfg *Config) {
	if port := v.GetInt("port"); port != 0 {
		cfg.Server.Port = port
	}

	if host := v.GetString("redis-host"); host != "" {
		cfg.Redis.Host = host
	}
}

func bindFlags(v *viper.Viper) {
	pflag.String("env", "dev", "environment")
	pflag.Int("port", 0, "server port")
	pflag.String("redis-host", "", "redis host")
	pflag.Parse()

	v.BindPFlags(pflag.CommandLine)
}

func bindEnv(v *viper.Viper) {
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
