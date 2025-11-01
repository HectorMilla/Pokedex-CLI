package pokecache

import (
	"testing"
)

func TestCacheBasic(t *testing.T) {
	c := Cache{cache: make(map[string]cacheEntry)}
	key := "test-url"
	var val []byte
	c.Add(key, val)

	// Test retrieval
	got, ok := c.Get(key)
	if !ok {
		t.Errorf("Expected key %q to be present in cache", key)
	}
	if len(got) != len(val) {
		t.Errorf("Expected value %v, got %v", len(val), len(got))
	}

	// Test missing key
	_, ok = c.Get("missing")
	if ok {
		t.Errorf("Expected missing key to not be present")
	}
}
