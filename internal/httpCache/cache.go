package httpCache

import (
	"sync"
	"time"
)

type Cache struct {
	data   map[string]cacheEntry
	config cacheConfig
	mu     sync.RWMutex
}

type cacheConfig struct {
	expirationInterval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		data: make(map[string]cacheEntry),
		config: cacheConfig{
			expirationInterval: interval,
		},
	}

	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			cache.reapLoop()
		}
	}()

	return cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		value:     value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.data[key]
	if !exists {
		return nil, false
	}
	return entry.value, true
}

func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, entry := range c.data {
		if now.Sub(entry.createdAt) > c.config.expirationInterval {
			delete(c.data, key)
		}
	}
}
