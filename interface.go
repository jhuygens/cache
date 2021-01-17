package cache

// Cache interface
type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, v interface{}) error
	Del(key string) error
	Expire(key string) error
	Clear() error
}
