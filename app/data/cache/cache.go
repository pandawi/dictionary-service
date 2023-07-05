package cache

import (
	"sync"
	"time"
)

type Cache struct {
	items      map[string]Item
	mutex      sync.RWMutex
	expiration time.Duration
}

type Item struct {
	value      []string
	expiration time.Time
}

func NewCache(expiration time.Duration) *Cache {
	cache := &Cache{
		items:      make(map[string]Item),
		expiration: expiration,
	}
	go cache.cleanupExpiredItems()
	return cache
}

func (c *Cache) Set(key string, value []string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items[key] = Item{
		value:      value,
		expiration: time.Now().Add(c.expiration),
	}
}

func (c *Cache) Get(key string) ([]string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	item, found := c.items[key]
	if !found {
		return nil, false
	}
	if time.Now().After(item.expiration) {
		return nil, false
	}
	return item.value, true
}

func (c *Cache) cleanupExpiredItems() {
	for {
		time.Sleep(c.expiration)
		c.mutex.Lock()
		for key, item := range c.items {
			if time.Now().After(item.expiration) {
				delete(c.items, key)
			}
		}
		c.mutex.Unlock()
	}
}
