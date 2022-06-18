package main

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	row0 := false
	colum0 := false

	// 判断第0列有没有0
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			colum0 = true
			break
		}
	}

	// 判断第0行有没有0
	for i := 0; i < n; i++ {
		if matrix[0][i] == 0 {
			row0 = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 这里只需要从第1行和第一列判断0就好
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if row0 {
		for i := 0; i < n; i++ {
			matrix[0][i] = 0
		}
	}
	if colum0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

}
