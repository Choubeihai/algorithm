package main

import (
	"fmt"
	"math"
)

// // 结合leetcode 55

func main() {
	var nums = []int{2, 3, 1, 1, 4}
	res := jump1(nums)
	fmt.Println(res)
}

// 贪心
func jump1(nums []int) int {
	n := len(nums)
	reach := make([]int, n) // 记录该点可以到达的最远的index

	var res int

	for i := 0; i < n; i++ {
		reach[i] = i + nums[i]
	}

	index := 0
	for index < n-1 {
		res++
		if reach[index] >= n-1 {
			break
		}
		canReach := reach[index]
		maxReach := reach[index]
		index++
		for i := index; i <= canReach; i++ {
			if reach[i] > maxReach {
				index = i
				maxReach = reach[i]
			}
		}
	}

	return res
}

// 动态规划
func jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)    // dp[i] 表示到当前index i的最小跳数
	reach := make([]int, n) // 记录该点可以到达的最远的index

	for i := 0; i < n; i++ {
		reach[i] = i + nums[i]
	}

	// dp数组初始化
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if reach[j] >= i {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
