package utilities

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

func CacheInit(defaultExpiration, cleanupInterval time.Duration) {
	c = cache.New(defaultExpiration, cleanupInterval)
}

func SetInCache(key string, value interface{}, expiration time.Duration) {
	if c == nil {
		panic("cache not initialized")
	}
	c.Set(key, value, expiration)
}

func GetFromCache(key string) (interface{}, bool) {
	if c == nil {
		panic("cache not initialized")
	}
	return c.Get(key)
}
