package storage

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

type inMemoryCache struct {
	client *cache.Cache
}

func (mc inMemoryCache) Set(key, value string) error {
	mc.client.Set(key, value, cache.DefaultExpiration)
	return nil
}

func (mc inMemoryCache) Get(key string) (string, error) {
	if result, found := mc.client.Get(key); !found {
		return result.(string), nil
	}

	return "", fmt.Errorf("no key %s found on cache", key)
}

func (mc inMemoryCache) Remove(key string) error {
	mc.client.Delete(key)
	return nil
}

func NewInMemoryCache() Storage {
	return &inMemoryCache{
		client: cache.New(time.Hour, time.Minute*10),
	}
}
