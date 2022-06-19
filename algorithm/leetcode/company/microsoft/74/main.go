package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	x := 0
	y := 0
	for x < m && y < n {
		if matrix[x][y] == target {
			return true
		} else if x+1 < m && matrix[x+1][y] == target {
			return true
		} else if x+1 < m && matrix[x+1][y] < target {
			x++
		} else if y+1 < n && matrix[x][y+1] == target {
			return true
		} else if y+1 < n && matrix[x][y+1] < target {
			y++
		} else {
			return false
		}
	}
	return false
}
