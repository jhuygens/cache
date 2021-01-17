package cache

import "fmt"

var cache Cache

// New doc ...
func New(c Cache) *Cache {
	return &c
}

// Register doc ...
func Register(c Cache) {
	cache = c
}

// ValidateRegisterImplement doc ...
func ValidateRegisterImplement() error {
	if cache != nil {
		return nil
	}
	return fmt.Errorf("The implementation of the 'Cache' interface has not been registered")
}

// Get doc ...
func Get(key string) (string, error) {
	err := ValidateRegisterImplement()
	if err != nil {
		return "", err
	}
	return cache.Get(key)
}

// Set doc ...
func Set(key string, s string) error {
	err := ValidateRegisterImplement()
	if err != nil {
		return err
	}
	return cache.Set(key, s)
}

// Del doc
func Del(key string) error {
	err := ValidateRegisterImplement()
	if err != nil {
		return err
	}
	return cache.Del(key)
}

// Expire cache
// sec: time to expire cache in seconds
func Expire(key string, sec int) error {
	err := ValidateRegisterImplement()
	if err != nil {
		return err
	}
	return cache.Expire(key, sec)
}

// TTL returns expire time in seconds
func TTL(key string) (int, error) {
	err := ValidateRegisterImplement()
	if err != nil {
		return 0, err
	}
	return cache.TTL(key)
}

// Close connection ...
func Close() error {
	err := ValidateRegisterImplement()
	if err != nil {
		return err
	}
	return cache.Close()
}
