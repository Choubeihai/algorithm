package main

func spiralOrder(matrix [][]int) []int {
	var m = len(matrix)
	var n = len(matrix[0])

	var left = 0
	var right = n - 1
	var top = 0
	var bottom = m - 1

	var res []int

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		for i := top + 1; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		if left < right && top < bottom {
			for i := right - 1; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}

			for i := bottom - 1; i > top; i-- {
				res = append(res, matrix[i][left])
			}
		}

		left++
		right--
		top++
		bottom--
	}
	return res
}
