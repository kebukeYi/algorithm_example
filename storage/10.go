package storage

import (
	"container/heap"
	"sync"
	"time"
)

type DelayItem struct {
	Value string
	At    time.Time
}

type DelayHeap []DelayItem

func (h DelayHeap) Len() int {
	return len(h)
}

func (h DelayHeap) Less(i, j int) bool {
	return h[i].At.Before(h[j].At)
}

func (h DelayHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *DelayHeap) Push(x any) {
	*h = append(*h, x.(DelayItem))
}

func (h *DelayHeap) Pop() any {
	old := *h
	item := old[len(old)-1]
	*h = old[:len(old)-1]
	return item
}

type DelayQueue struct {
	mu sync.Mutex
	h  DelayHeap
}

func (q *DelayQueue) Add(item DelayItem) {
	q.mu.Lock()
	defer q.mu.Unlock()

	heap.Push(&q.h, item)
}

func (q *DelayQueue) PopReady(now time.Time) (DelayItem, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.h) == 0 || q.h[0].At.After(now) {
		return DelayItem{}, false
	}

	return heap.Pop(&q.h).(DelayItem), true
}
