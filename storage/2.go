package storage

import "container/heap"

func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func SearchRotated(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 返回最大的 k 个数
func TopK(nums []int, k int) []int {
	if k <= 0 {
		return nil
	}

	h := &IntMinHeap{}
	heap.Init(h)

	for _, x := range nums {
		heap.Push(h, x)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ans := make([]int, h.Len())
	for i := len(ans) - 1; i >= 0; i-- {
		ans[i] = heap.Pop(h).(int)
	}
	return ans
}

func DailyTemperatures(nums []int) []int {
	ans := make([]int, len(nums))
	stack := []int{}

	for i, temperature := range nums {
		for len(stack) > 0 &&
			nums[stack[len(stack)-1]] < temperature {

			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[j] = i - j
		}
		stack = append(stack, i)
	}
	return ans
}

func BFS(graph map[int][]int, start int) []int {
	visited := map[int]bool{start: true}
	queue := []int{start}
	ans := []int{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		ans = append(ans, node)

		for _, next := range graph[node] {
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}
	return ans
}

func DFS(graph map[int][]int, start int) []int {
	visited := map[int]bool{}
	ans := []int{}

	var visit func(int)
	visit = func(node int) {
		if visited[node] {
			return
		}
		visited[node] = true
		ans = append(ans, node)

		for _, next := range graph[node] {
			visit(next)
		}
	}

	visit(start)
	return ans
}

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)

	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{parent: parent, size: size}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(a, b int) bool {
	ra, rb := d.Find(a), d.Find(b)
	if ra == rb {
		return false
	}

	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}

	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
	return true
}
