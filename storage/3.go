package storage

import "sort"

// 0/1 背包
func Knapsack(weights, values []int, capacity int) int {
	dp := make([]int, capacity+1)

	for i, weight := range weights {
		for c := capacity; c >= weight; c-- {
			dp[c] = max(dp[c], dp[c-weight]+values[i])
		}
	}
	return dp[capacity]
}

// 最长递增子序列
func LengthOfLIS(nums []int) int {
	tails := []int{}

	for _, x := range nums {
		i := sort.SearchInts(tails, x)

		if i == len(tails) {
			tails = append(tails, x)
		} else {
			tails[i] = x
		}
	}
	return len(tails)
}

// 网格最小路径和
func MinPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	dp := make([]int, len(grid[0]))

	for i := range grid {
		for j := range grid[0] {
			if i == 0 && j == 0 {
				dp[j] = grid[i][j]
			} else if i == 0 {
				dp[j] = dp[j-1] + grid[i][j]
			} else if j == 0 {
				dp[j] += grid[i][j]
			} else {
				dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
			}
		}
	}
	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
