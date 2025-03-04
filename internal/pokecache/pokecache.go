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
	mu    sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	var c = Cache{
		entry: map[string]cacheEntry{},
		mu:    sync.Mutex{},
	}
	go c.reapLoop(interval)
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	val, ok := c.entry[key]
	if !ok {
		return []byte{}, false
	}
	c.mu.Unlock()
	return val.value, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	tic := time.NewTicker(interval)
	defer tic.Stop()

	for {
		t, ok := <-tic.C
		if !ok {
			return
		}
		c.mu.Lock()

		// fmt.Println("ticker: ", t.String())
		checkTime := t.Add(-interval)
		// fmt.Println("ticker - interval: ", checkTime.String())
		for k, v := range c.entry {
			// fmt.Println("created at: ", v.createdAt.String())
			if v.createdAt.Before(checkTime) {
				delete(c.entry, k)
			}
		}
		c.mu.Unlock()
	}
}
