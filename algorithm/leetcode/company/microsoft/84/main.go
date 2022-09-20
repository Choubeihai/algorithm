package main

func main() {

}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 1 {
		return heights[0]
	}
	stack := make([]int, n) // 单调递增栈
	p := -1
	res := 0
	for i := 0; i < n; i++ {
		for p >= 0 && heights[stack[p]] >= heights[i] {
			cur := stack[p]
			p--
			left := 0
			right := 0
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
		left := 0
		right := n - 1
		if p == -1 {
			left = 0
		} else {
			left = stack[p] + 1
		}
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
