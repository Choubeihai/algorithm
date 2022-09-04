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
			index := stack[p]
			height := heights[index]
			p--
			var left int
			var right int
			if p == -1 {
				left = 0
			} else {
				left = stack[p]
			}
			right = i - 1
			res = max(res, (right-left+1)*height)
		}
		p++
		stack[p] = i
	}

	for p >= 0 {
		index := stack[p]
		height := heights[index]
		p--
		if p == -1 {

		}

		var width int
		if p == -1 {
			width = n
		} else {
			width = (n - 1) - (stack[p] + 1) - 1
		}
		res = max(res, width*height)
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
