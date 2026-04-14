package redis

// Deinfe abstract distributed cache interface
type IDistCache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
}
