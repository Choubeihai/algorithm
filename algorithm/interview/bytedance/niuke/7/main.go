package main

import (
	"math"
)

/**
一个N*M的的数组里面有0和1, 1代表障碍，求从左上角到右下角最短路径，并打印。可以上下左右走
*/

var res int

func findShortestPath(matrix [][]int) int {
	res = math.MaxInt32
	dfs(matrix, 0, 0, 0)
	return res
}

func dfs(matrix [][]int, x, y int, sum int) {
	if x == len(matrix) && y == len(matrix[0]) {
		res = min(res, sum)
		return
	}
	if x >= len(matrix) || y >= len(matrix[0]) {
		return
	}
	if matrix[x][y] == 1 {
		return
	}
	dfs(matrix, x+1, y, sum+1)
	dfs(matrix, x, y+1, sum+1)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
