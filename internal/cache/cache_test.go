package cache

import (
	gocache "github.com/patrickmn/go-cache"
	"testing"
)

func TestLoadInMemory(t *testing.T) {
	tests := []struct {
		name           string
		usedCacheId    string
		queriedCacheId string
		want           bool
	}{
		{"retrieve_cache_success", "cacheId", "cacheId", true},
		{"retrieve_cache_fail", "cacheId", "wrongCacheId", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadInMemory(); &got != nil {
				_ = got.Set(tt.usedCacheId, "value", gocache.DefaultExpiration)
				if _, found := got.Get(tt.queriedCacheId); found != tt.want {
					t.Errorf("LoadInMemory() = %v, want %v", found, tt.want)
				}
			} else {
				t.Errorf("Couldn't LoadInMemory()")
			}
		})
	}
}

func TestManager_LoadInMemory(t *testing.T) {
	var c = LoadInMemory()
	c.Set("name", "test", 0)
	c.Set("age", 55, 0)

	if _, found := c.Get("name"); !found {
		t.Errorf("Couldn't find item")
	}

	if _, found := c.Get("age"); !found {
		t.Errorf("Couldn't find item")
	}

	var c1 = LoadInMemory()
	c1.Set("name2", "test2", 0)
	c1.Set("name3", "test3", 0)

	if _, found := c1.Get("name2"); !found {
		t.Errorf("Couldn't find item")
	}

	if _, found := c1.Get("name3"); !found {
		t.Errorf("Couldn't find item")
	}

	if _, found := c1.Get("name"); !found {
		t.Errorf("Couldn't find item")
	}

	if c != c1 {
		t.Errorf("caches inMemory addresses should be same")
	}
}
