package storage

import (
	"hash/crc32"
	"sort"
	"sync"
)

type ConsistentHash struct {
	mu     sync.RWMutex
	vnodes int
	keys   []uint32
	nodes  map[uint32]string
}

func NewConsistentHash(vnodes int) *ConsistentHash {
	return &ConsistentHash{
		vnodes: vnodes,
		nodes:  make(map[uint32]string),
	}
}

func (c *ConsistentHash) Add(node string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i := 0; i < c.vnodes; i++ {
		key := crc32.ChecksumIEEE(
			[]byte(node + "#" + string(rune(i))),
		)
		c.keys = append(c.keys, key)
		c.nodes[key] = node
	}

	sort.Slice(c.keys, func(i, j int) bool {
		return c.keys[i] < c.keys[j]
	})
}

func (c *ConsistentHash) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if len(c.keys) == 0 {
		return "", false
	}

	hash := crc32.ChecksumIEEE([]byte(key))
	index := sort.Search(len(c.keys), func(i int) bool {
		return c.keys[i] >= hash
	})

	if index == len(c.keys) {
		index = 0
	}

	return c.nodes[c.keys[index]], true
}
