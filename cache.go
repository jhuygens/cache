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
func Get(key string) (interface{}, error) {
	err := ValidateRegisterImplement()
	if err != nil {
		return nil, err
	}
	return cache.Get(key)
}

// Set doc ...
func Set(key string, v interface{}) error {
	err := ValidateRegisterImplement()
	if err != nil {
		return err
	}
	return cache.Set(key, v)
}

// Del doc
func Del(key string) error {
	err := ValidateRegisterImplement()
	if err != nil {
		return err
	}
	return cache.Del(key)
}

// Expire doc ...
func Expire(key string) error {
	err := ValidateRegisterImplement()
	if err != nil {
		return err
	}
	return cache.Expire(key)
}

// Clear doc ...
func Clear() error {
	err := ValidateRegisterImplement()
	if err != nil {
		return err
	}
	return cache.Clear()
}
