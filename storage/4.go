package storage

import (
	"container/list"
	"sync"
)

type lruEntry struct {
	key   int
	value int
}

type LRUCache struct {
	mu    sync.Mutex
	cap   int
	items map[int]*list.Element
	list  *list.List
}

func NewLRU(capacity int) *LRUCache {
	return &LRUCache{
		cap:   capacity,
		items: make(map[int]*list.Element),
		list:  list.New(),
	}
}

func (c *LRUCache) Get(key int) (int, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	elem, ok := c.items[key]
	if !ok {
		return 0, false
	}

	c.list.MoveToFront(elem)
	return elem.Value.(lruEntry).value, true
}

func (c *LRUCache) Put(key, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.items[key]; ok {
		elem.Value = lruEntry{key, value}
		c.list.MoveToFront(elem)
		return
	}

	c.items[key] = c.list.PushFront(lruEntry{key, value})

	if c.list.Len() > c.cap {
		elem := c.list.Back()
		delete(c.items, elem.Value.(lruEntry).key)
		c.list.Remove(elem)
	}
}
