package main

/**
N皇后问题
*/

var res [][]string

func solveNQueens(n int) [][]string {
	res = nil
	var board = make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
	}
	dfs(board, 0)
	return res
}

func dfs(board [][]byte, row int) {
	if row == len(board) {
		var tmpRes []string
		for i := 0; i < len(board); i++ {
			s := ""
			for j := 0; j < len(board[0]); j++ {
				if board[i][j] == 'Q' {
					s += string(board[i][j])
				} else {
					s += "."
				}
			}
			tmpRes = append(tmpRes, s)
		}
		res = append(res, tmpRes)
		return
	}

	// 这一行的所有列
	for i := 0; i < len(board[0]); i++ {
		x := row
		y := i
		flag := false // 没有Q

		// 正对角线
		x--
		y--
		for x >= 0 && y >= 0 {
			if board[x][y] == 'Q' {
				flag = true
				break
			}
			x--
			y--
		}
		// 斜对角线
		if !flag {
			x = row
			y = i
			x--
			y++
			for x >= 0 && y < len(board[0]) {
				if board[x][y] == 'Q' {
					flag = true
					break
				}
				x--
				y++
			}
		}
		// 同行
		if !flag {
			x = row
			y = i
			y--
			for y >= 0 {
				if board[x][y] == 'Q' {
					flag = true
					break
				}
				y--
			}
		}
		// 同列
		if !flag {
			x = row
			y = i
			x--
			for x >= 0 {
				if board[x][y] == 'Q' {
					flag = true
					break
				}
				x--
			}
		}

		if flag {
			continue
		} else {
			board[row][i] = 'Q'
			dfs(board, row+1)
			board[row][i] = '0'
		}
	}
}
