package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	var i = m - 1
	var j = 0

	for i >= 0 && j <= n-1 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			i--
		} else {
			j++
		}
	}
	return false
}
