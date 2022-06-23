package main

func maximalRectangle(matrix [][]byte) int {
	var res = 0
	var heights = make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				heights[j] += 1
			} else {
				heights[j] = 0
			}
		}
		ans := largestRectangleArea(heights)
		res = max(ans, res)
	}
	return res
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
	var stack = make([]int, n)
	var p = -1
	for i := 0; i < n; i++ {
		for p != -1 && heights[stack[p]] > heights[i] {
			index := stack[p]
			height := heights[index]
			p--
			for p != -1 && heights[stack[p]] == height {
				p--
			}
			var width = 0
			if p == -1 {
				width = i
			} else {
				width = i - stack[p] - 1
			}
			res = max(res, width*height)
		}
		p++
		stack[p] = i
	}

	for p != -1 {
		index := stack[p]
		height := heights[index]
		p--
		for p != -1 && heights[stack[p]] == height {
			p--
		}
		var width = 0
		if p == -1 {
			width = n
		} else {
			width = n - stack[p] - 1
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
