package storage

import "math/rand"

const maxLevel = 16

type skipNode struct {
	key, value int
	next       []*skipNode
}

type SkipList struct {
	head  *skipNode
	level int
}

func NewSkipList() *SkipList {
	return &SkipList{
		head:  &skipNode{next: make([]*skipNode, maxLevel)},
		level: 1,
	}
}

func randomLevel() int {
	level := 1
	for level < maxLevel && rand.Intn(4) == 0 {
		level++
	}
	return level
}

func (s *SkipList) Get(key int) (int, bool) {
	cur := s.head

	for level := s.level - 1; level >= 0; level-- {
		for cur.next[level] != nil &&
			cur.next[level].key < key {
			cur = cur.next[level]
		}
	}

	cur = cur.next[0]
	if cur != nil && cur.key == key {
		return cur.value, true
	}
	return 0, false
}

func (s *SkipList) Set(key, value int) {
	update := make([]*skipNode, maxLevel)
	cur := s.head

	for level := s.level - 1; level >= 0; level-- {
		for cur.next[level] != nil &&
			cur.next[level].key < key {
			cur = cur.next[level]
		}
		update[level] = cur
	}

	cur = cur.next[0]
	if cur != nil && cur.key == key {
		cur.value = value
		return
	}

	level := randomLevel()
	if level > s.level {
		for i := s.level; i < level; i++ {
			update[i] = s.head
		}
		s.level = level
	}

	node := &skipNode{
		key:   key,
		value: value,
		next:  make([]*skipNode, level),
	}

	for i := 0; i < level; i++ {
		node.next[i] = update[i].next[i]
		update[i].next[i] = node
	}
}
