package main

// 结合leetcode 45

func canJump(nums []int) bool {
	var n = len(nums)
	var reach = make([]int, n) // 从此处能达到的最远index
	for i := 0; i < n; i++ {
		reach[i] = i + nums[i]
	}

	canMaxReach := 0 // 选举出最大能到达的位置

	for i := 0; i < n; i++ {
		if i <= canMaxReach {
			canMaxReach = max(canMaxReach, reach[i])
			if canMaxReach >= n-1 {
				return true
			}
		} else { // 能到达的最远index小于当前位置，说明达到不了此处，直接退出循环
			break
		}
	}

	return false
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
