package ristretto

// Define abstract in-memory cache interface
type IInMemCache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
}
