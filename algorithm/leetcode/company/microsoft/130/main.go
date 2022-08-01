package main

func solve(board [][]byte) {
	var m = len(board)
	var n = len(board[0])
	if m == 0 || n == 0 {
		return
	}
	for i := 0; i < m; i++ {
		dfs(board, i, 0)
		dfs(board, i, n-1)
	}
	for i := 0; i < n; i++ {
		dfs(board, 0, i)
		dfs(board, m-1, i)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	dfs(board, x, y-1)
	dfs(board, x, y+1)
	dfs(board, x-1, y)
	dfs(board, x+1, y)
}
