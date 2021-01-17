package cache

// Cache interface
type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, v interface{}) error
	Del(key string) error
	SetExpire(key string, sec int) error
	GetExpireTime(key string) (int, error)
	Clear() error
}
