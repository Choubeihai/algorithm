package main

/**
思路：
S = (y-x)*min{height(x), height(y)}
这就简单了
*/

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	res := 0

	for left < right {
		h := min(height[left], height[right])
		res = max(res, h*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
