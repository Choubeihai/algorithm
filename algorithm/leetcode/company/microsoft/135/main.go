package main

func candy(ratings []int) int {
	var n = len(ratings)
	var left = make([]int, n)
	var right = make([]int, n)
	var res int

	// 比左
	for i := 0; i < n; i++ {
		if i == 0 {
			left[i] = 1
		} else {
			if ratings[i] > ratings[i-1] {
				left[i] = left[i-1] + 1
			} else {
				left[i] = 1
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			right[i] = 1
		} else {
			if ratings[i] > ratings[i+1] {
				right[i] = right[i+1] + 1
			} else {
				right[i] = 1
			}
		}
		res += max(left[i], right[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
