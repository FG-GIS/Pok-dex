package pokecache

import (
	// "fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	entry map[string]cacheEntry
	mu    *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	var c = Cache{
		entry: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.entry[key]
	if !ok {
		return []byte{}, false
	}
	return val.value, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	tic := time.NewTicker(interval)
	// defer tic.Stop()

	for range tic.C {
		t := time.Now()

		func() {
			c.mu.Lock()
			defer c.mu.Unlock()

			checkTime := t.Add(-interval)
			for k, v := range c.entry {
				if v.createdAt.Before(checkTime) {
					delete(c.entry, k)
				}
			}
		}()
	}
}
