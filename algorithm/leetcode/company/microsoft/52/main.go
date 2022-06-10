package main

/**
一般来说，求方案个数用动态规划，列举方案用回溯
但是这个求方案个数无法使用动态规划，找不到好的状态表示。
因此仍然使用回溯，利用51的思路。
*/

var nn int
var res int

func totalNQueens(n int) int {
	res = 0
	nn = n
	var board = make([][]byte, nn)
	for i := 0; i < len(board); i++ {
		board[i] = make([]byte, nn)
	}
	dfs(board, 0)
	return res
}

func dfs(board [][]byte, row int) {
	if row == nn {
		res++
		return
	}
	for i := 0; i < nn; i++ {
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
			for x >= 0 && y < nn {
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
			continue // 有Q，那就直接下一列
		} else {
			board[row][i] = 'Q'
			dfs(board, row+1)
			board[row][i] = '0'
		}
	}
}
