package main

func main() {

}

// 和84题很像
func trap(height []int) int {
	var n = len(height)
	if n == 1 || n == 2 {
		return 0
	}
	var res int
	var stack = make([]int, n) // 严格单调递减栈
	var p = -1

	for i := 0; i < n; i++ {
		for p >= 0 && height[stack[p]] < height[i] {
			cur := stack[p]
			p--
			if p == -1 {
				break
			}
			h := min(height[stack[p]], height[i]) - height[cur]
			var left = stack[p] + 1
			var right = i - 1
			res += (right - left + 1) * h
		}
		p++
		stack[p] = i
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
