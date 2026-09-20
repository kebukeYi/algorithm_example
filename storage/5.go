package storage

import (
	"sync"
	"time"
)

type TTLItem[V any] struct {
	value   V
	expires time.Time
}

type TTLCache[V any] struct {
	mu    sync.RWMutex
	items map[string]TTLItem[V]
}

func NewTTLCache[V any]() *TTLCache[V] {
	return &TTLCache[V]{
		items: make(map[string]TTLItem[V]),
	}
}

func (c *TTLCache[V]) Set(key string, value V, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = TTLItem[V]{
		value:   value,
		expires: time.Now().Add(ttl),
	}
}

func (c *TTLCache[V]) Get(key string) (V, bool) {
	c.mu.RLock()
	item, ok := c.items[key]
	c.mu.RUnlock()

	var zero V
	if !ok || time.Now().After(item.expires) {
		if ok {
			c.mu.Lock()
			delete(c.items, key)
			c.mu.Unlock()
		}
		return zero, false
	}

	return item.value, true
}
