package main

func maxArea(height []int) int {
	var res = 0
	var left = 0
	var right = len(height) - 1
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > res {
			res = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}
