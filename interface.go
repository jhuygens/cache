package cache

// Cache interface
type Cache interface {
	Get(key string) (string, error)
	Set(key string, s string) error
	Del(key string) error
	SetExpire(key string, sec int) error
	GetExpireTime(key string) (int, error)
	Clear() error
}
