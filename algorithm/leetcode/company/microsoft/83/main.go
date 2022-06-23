package main

import "fmt"

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(heights))
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
		if p != -1 && heights[stack[p]] > heights[i] {
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
			res = max(res, height*width)
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
		res = max(res, height*width)
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
