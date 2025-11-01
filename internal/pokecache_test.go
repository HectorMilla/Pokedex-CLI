package pokeapi

import (
	"testing"
)

func TestCacheBasic(t *testing.T) {
	c := Cache{entries: make(map[string]cacheEntry)}
	key := "test-url"
	val := LocationArea{Count: 1}
	c.Add(key, val)

	// Test retrieval
	got, ok := c.Get(key)
	if !ok {
		t.Errorf("Expected key %q to be present in cache", key)
	}
	if got.Count != val.Count {
		t.Errorf("Expected value %v, got %v", val.Count, got.Count)
	}

	// Test missing key
	_, ok = c.Get("missing")
	if ok {
		t.Errorf("Expected missing key to not be present")
	}
}
