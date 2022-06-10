package main

import (
	"bytes"
)

/**
N皇后问题
*/

var nn int
var res [][]string

func solveNQueens(n int) [][]string {
	nn = n
	res = nil
	var board = make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
	}
	dfs(board, 0)
	return res
}

func dfs(board [][]byte, row int) {
	if row == nn {
		var tmpRes []string
		for i := 0; i < len(board); i++ {
			bb := &bytes.Buffer{}
			for j := 0; j < len(board[0]); j++ {
				if board[i][j] == 'Q' {
					bb.WriteByte(board[i][j])
				} else {
					bb.WriteByte('.')
				}
			}
			tmpRes = append(tmpRes, bb.String())
		}
		res = append(res, tmpRes)
		return
	}

	// 这一行的所有列
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
			continue
		} else {
			board[row][i] = 'Q'
			dfs(board, row+1)
			board[row][i] = '0'
		}
	}
}
