package cache

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheSetGet(t *testing.T) {
	c := NewMemoryCache()
	c.Set("India", "New Delhi")

	v, found := c.Get("India")

	assert.True(t, found)
	assert.Equal(t, "New Delhi", v)
}

func TestCacheMiss(t *testing.T) {
	c := NewMemoryCache()
	v, found := c.Get("USA")
	assert.False(t, found)
	assert.Nil(t, v)
}

func TestCacheConcurrentAccess(t *testing.T) {
	cache := NewMemoryCache()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := "country"
			cache.Set(key, i)
			cache.Get(key)
		}(i)
	}
	wg.Wait()
}
