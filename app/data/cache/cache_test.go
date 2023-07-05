package cache_test

import (
	"dictionary-service/app/data/cache"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCache_SetAndGet(t *testing.T) {
	cache := cache.NewCache(1 * time.Second)
	cache.Set("key1", []string{"value1"})
	value, found := cache.Get("key1")
	assert.True(t, found)
	assert.ElementsMatch(t, value, []string{"value1"})
}

func TestCache_GetExpiredItem(t *testing.T) {
	cache := cache.NewCache(2 * time.Second)
	cache.Set("key1", []string{"value1"})
	time.Sleep(3 * time.Second)
	_, found := cache.Get("key1")
	assert.False(t, found)
}

func TestCache_GetNonExistentItem(t *testing.T) {
	cache := cache.NewCache(2 * time.Second)
	cache.Set("key1", []string{"value1"})
	time.Sleep(3 * time.Second)
	_, found := cache.Get("key2")
	assert.False(t, found)
}
