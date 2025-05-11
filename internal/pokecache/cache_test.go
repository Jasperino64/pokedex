package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache(1)
	if cache == nil {
		t.Fatal("Expected non-nil cache")
	}

	key := "testKey"
	value := []byte("testValue")

	cache.Add(key, value)

	cachedValue, found := cache.Get(key)
	if !found {
		t.Fatal("Expected to find cached value")
	}
	if string(cachedValue) != string(value) {
		t.Fatalf("Expected %s, got %s", value, cachedValue)
	}

	time.Sleep(2 * time.Second)
	_, found = cache.Get(key)
	if found {
		t.Fatal("Expected to not find cached value after expiration")
	}
}
