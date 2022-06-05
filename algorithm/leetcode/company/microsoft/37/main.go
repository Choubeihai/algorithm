package main

func main() {

}

/**
思路：回溯的方法
首先建立起来3个map，标识在行，列，box里有哪些数字
然后进行回溯查找所有的方法。
*/

var rowMap [][]int
var columnMap [][]int
var boxMap [][][]int

func solveSudoku(board [][]byte) {
	rowMap = make([][]int, 9)
	for i := 0; i < 9; i++ {
		rowMap[i] = make([]int, 10)
	}
	columnMap = make([][]int, 9)
	for i := 0; i < 9; i++ {
		columnMap[i] = make([]int, 10)
	}
	boxMap = make([][][]int, 3)
	for i := 0; i < 3; i++ {
		boxMap[i] = make([][]int, 3)
		for j := 0; j < 3; j++ {
			boxMap[i][j] = make([]int, 10)
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '0'
				rowMap[i][num] = 1
				columnMap[j][num] = 1
				boxMap[i/3][j/3][num] = 1
			}
		}
	}
	dfs(board, 0, 0)
}

func dfs(board [][]byte, x, y int) bool {
	if y >= len(board[0]) {
		y = 0
		x = x + 1
	}
	if x == len(board) {
		return true
	}
	if board[x][y] == '.' {
		// 填充数字
		for i := 1; i <= 9; i++ {
			if rowMap[x][i] != 1 && columnMap[y][i] != 1 && boxMap[x/3][y/3][i] != 1 {
				rowMap[x][i] = 1
				columnMap[y][i] = 1
				boxMap[x/3][y/3][i] = 1
				board[x][y] = byte(i) + '0'
				if dfs(board, x, y+1) {
					return true
				}
				rowMap[x][i] = 0
				columnMap[y][i] = 0
				boxMap[x/3][y/3][i] = 0
				board[x][y] = '.'
			}
		}
	} else {
		return dfs(board, x, y+1)
	}
	return false
}
