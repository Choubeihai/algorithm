package main

func main() {
}

// 思路：把该行是否出现过1-9做一个标记，把该列是否出现过1-9做一个标记，把该box内是否出现过1-9做一个标记
func isValidSudoku(board [][]byte) bool {
	var rowsMap = make([][]int, 9)
	for i := 0; i < 9; i++ {
		rowsMap[i] = make([]int, 10) // 标识本行出现1-9中哪些数字
	}
	var columnMap = make([][]int, 9)
	for i := 0; i < 9; i++ {
		columnMap[i] = make([]int, 10) // 标识本列出现1-9中哪些数字

	}
	var boxMap = make([][][]int, 3) // 把boxMap = make([][][]int, 3) // [3][3][10]
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
				if boxMap[i/3][j/3][num] == 1 {
					return false
				} else {
					boxMap[i/3][j/3][num] = 1
				}

				if rowsMap[i][num] == 1 {
					return false
				} else {
					rowsMap[i][num] = 1
				}
				if columnMap[j][num] == 1 {
					return false
				} else {
					columnMap[j][num] = 1
				}
			}

		}
	}
	return true
}
