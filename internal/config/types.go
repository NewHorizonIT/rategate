package config

type Config struct {
	App       AppConfig       `mapstructure:"app"`
	Server    ServerConfig    `mapstructure:"server"`
	Redis     RedisConfig     `mapstructure:"redis"`
	RateLimit RateLimitConfig `mapstructure:"rateLimit"`
	Database  DatabaseConfig  `mapstructure:"database"`
	Cache     CacheConfig     `mapstructure:"cache"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
}

type RateLimitConfig struct {
	Strategy string `mapstructure:"strategy"`
	Requests int    `mapstructure:"requests"`
	Window   int    `mapstructure:"window"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
}

// Structure for Ristretto cache configuration
type CacheConfig struct {
	NumCounters int `mapstructure:"numCounters"`
	MaxCost     int `mapstructure:"maxCost"`
	BufferItems int `mapstructure:"bufferItems"`
	TTL         int `mapstructure:"ttl"` // Time to live in milliseconds
}
