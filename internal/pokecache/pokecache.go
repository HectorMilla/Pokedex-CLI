package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mux   *sync.Mutex
	cache map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	ticker := time.NewTicker(interval)
	go c.reapLoop(ticker.C)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	newEntry := cacheEntry{time.Now(), val}
	c.mux.Lock()
	c.cache[key] = newEntry
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	entry, exist := c.cache[key]
	c.mux.Unlock()
	if !exist {
		return []byte{}, exist
	}
	return entry.val, exist
}

func (c *Cache) reapLoop(time <-chan time.Time) {
	minTime := <-time
	c.mux.Lock()
	for entry := range c.cache {
		if c.cache[entry].createdAt.Compare(minTime) == -1 {
			delete(c.cache, entry)
		}
	}
	c.mux.Unlock()
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
