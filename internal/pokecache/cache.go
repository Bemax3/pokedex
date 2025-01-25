package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	Data  map[string]cacheEntry
	mu    *sync.Mutex
	inter time.Duration
}

func NewCache(inter time.Duration) *Cache {
	cache := &Cache{
		Data:  make(map[string]cacheEntry, 10),
		mu:    &sync.Mutex{},
		inter: inter,
	}

	defer cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.Data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	cacheEntry, exist := c.Data[key]
	c.mu.Unlock()
	return cacheEntry.val, exist
}

func (c *Cache) reapLoop() {

	timer := time.NewTicker(c.inter)

	go func() {
		for {
			select {
			case <-timer.C:
				c.mu.Lock()
				for k, v := range c.Data {
					if v.createdAt.Add(c.inter).Before(time.Now()) {
						delete(c.Data, k)
					}
				}
				c.mu.Unlock()
			}
		}
	}()
}
