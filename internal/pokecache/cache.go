package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entry map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Entry: make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.Entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	d, ok := c.Entry[key]

	if !ok {
		return []byte{}, false
	}

	return d.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	// We need "go" here to run this in the background
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		// Range over the map entries
		for key, entry := range c.Entry {
			// Check if this entry is too old
			if now.Sub(entry.createdAt) > interval {
				// Delete old entries
				delete(c.Entry, key)
			}
		}
		c.mu.Unlock()
	}
}
