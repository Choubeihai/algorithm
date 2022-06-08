package main

func main() {

}

// 动态规划 前缀数组 后缀数组
func trap(height []int) int {
	var n = len(height)
	var leftMax = make([]int, n)  // leftMax[i] 表示index为i的左侧的最大值（比较中不包括index i）
	var rightMax = make([]int, n) // leftMax[i] 表示index为i的右侧的最大值（比较中不包括index i）
	var res int

	for i := 0; i < n; i++ {
		if i == 0 {
			leftMax[i] = 0
		} else {
			leftMax[i] = max(leftMax[i-1], height[i-1])
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			rightMax[i] = 0
		} else {
			rightMax[i] = max(rightMax[i+1], height[i+1])
		}
	}

	for i := 0; i < n; i++ {
		minMax := min(leftMax[i], rightMax[i])
		if minMax > height[i] {
			remain := minMax - height[i]
			res = res + remain
		}

	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
