package main

import "fmt"

func main() {
	var matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix)
	rotate(matrix)

}

func rotate(matrix [][]int) {
	var m = len(matrix)
	var n = len(matrix[0])

	/**
	反转
	1 2 3     1 4 7
	4 4 6 ——> 2 4 8
	7 8 9     3 6 9
	*/
	for i := 0; i < m; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 沿着列中线 | 反转
	for i := 0; i < m; i++ {
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
	fmt.Println(matrix)
}
