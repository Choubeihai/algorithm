package main

func main() {

}

func largestRectangleArea(heights []int) int {
	var n = len(heights)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}
	var res = 0
	var stack = make([]int, n) // 严格单调递增栈
	var p = -1

	for i := 0; i < n; i++ {
		for p >= 0 && heights[stack[p]] >= heights[i] {
			cur := stack[p]
			p--
			var left int
			var right int
			if p == -1 {
				left = 0
			} else {
				left = stack[p] + 1
			}
			right = i - 1
			res = max(res, (right-left+1)*heights[cur])
		}
		p++
		stack[p] = i
	}

	for p >= 0 {
		cur := stack[p]
		p--
		var left int
		var right int
		if p == -1 {
			left = 0
		} else {
			left = stack[p] + 1
		}
		right = n - 1
		res = max(res, (right-left+1)*heights[cur])
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
