package cache

import "sync"

type memoryCache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewMemoryCache() Cache {
	return &memoryCache{
		data: make(map[string]interface{}),
	}
}

func (m *memoryCache) Get(key string) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	val, ok := m.data[key]
	return val, ok
}

func (m *memoryCache) Set(key string, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}
