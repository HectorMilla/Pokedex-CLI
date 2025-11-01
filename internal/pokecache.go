package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mux   *sync.Mutex
	cache map[string]cacheEntry
}

func NewCache(interval time.Duration, val LocationArea, url string) {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	ticker := time.NewTicker(interval)
	c.reapLoop(ticker.C)
	return c
}

func (cache *Cache) Add(key string, val LocationArea) {
	newEntry := cacheEntry{time.Now(), val}
	c.mux.Lock()
	c.cache[key] = newEntry
	c.mux.Unlock()
}

func (cache *Cache) Get(key string) (LocationArea, bool) {
	c.mux.Lock()
	entry, exist := c.cache[key]
	c.mux.Unlock()
	if !exist {
		return LocationArea{}, exist
	}
	return entry.val, exist
}

func (cache *Cache) reapLoop(time <-chan time.Time) {
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
	val       LocationArea
}
