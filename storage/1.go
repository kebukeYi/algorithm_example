package storage

import (
	"sort"
)

func TwoSum(nums []int, target int) []int {
	seen := map[int]int{}
	for i, x := range nums {
		if j, ok := seen[target-x]; ok {
			return []int{j, i}
		}
		seen[x] = i
	}
	return nil
}

// 无重复字符的最长子串
func LongestSubstring(s string) int {
	last := map[byte]int{}
	left, ans := 0, 0

	for right := 0; right < len(s); right++ {
		if p, ok := last[s[right]]; ok && p >= left {
			left = p + 1
		}
		last[s[right]] = right
		ans = max(ans, right-left+1)
	}
	return ans
}

// 和为 k 的连续子数组数量
func SubarraySum(nums []int, k int) int {
	count, prefix := 0, 0
	freq := map[int]int{0: 1}

	for _, x := range nums {
		prefix += x
		count += freq[prefix-k]
		freq[prefix]++
	}
	return count
}

func MergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0)
	for _, cur := range intervals {
		if len(ans) == 0 || ans[len(ans)-1][1] < cur[0] {
			ans = append(ans, []int{cur[0], cur[1]})
		} else {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], cur[1])
		}
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}

func HasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}

func MergeTwoLists(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next = a
			a = a.Next
		} else {
			tail.Next = b
			b = b.Next
		}
		tail = tail.Next
	}

	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return dummy.Next
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)
		level := make([]int, 0, n)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, level)
	}
	return ans
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(MaxDepth(root.Left), MaxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
