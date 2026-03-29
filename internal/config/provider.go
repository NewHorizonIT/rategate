package config

type Provider interface {
	Load() (*Config, error)
}
