package custom_cache

import (
	"errors"
)

type Cache struct {
	cache map[string]interface{}
}

func New() *Cache {
	return &Cache{
		cache: make(map[string]interface{}),
	}
}

func (c Cache) Set(key string, value interface{}) {
	c.cache[key] = value
}

func (c Cache) keyExist(key string) bool {
	_, exists := c.cache[key]
	return exists
}

func (c Cache) Get(key string) (interface{}, error) {
	exists := c.keyExist(key)
	if !exists {
		return nil, errors.New("no values for the key " + key + ".")
	}
	return c.cache[key], nil
}

func (c Cache) Delete(key string) error {
	exists := c.keyExist(key)
	if !exists {
		return errors.New("no values for the key " + key + ".")
	}
	delete(c.cache, key)
	return nil
}
