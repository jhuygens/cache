package cache

// Cache interface
type Cache interface {
	Get(key string) (string, error)
	Set(key string, s string) error
	Del(key string) error
	Expire(key string, sec int) error
	TTL(key string) (int, error)
	Close() error
}
