package main

var neighborWithEdgeO [][]bool

func solve(board [][]byte) {
	var m = len(board)
	var n = len(board[0])
	if m == 0 || n == 0 {
		return
	}
	neighborWithEdgeO = make([][]bool, m)
	for i := 0; i < m; i++ {
		neighborWithEdgeO[i] = make([]bool, n)
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
			if board[i][j] == 'O' && !neighborWithEdgeO[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return
	}
	if board[x][y] != 'O' {
		return
	}
	if neighborWithEdgeO[x][y] {
		return
	}
	neighborWithEdgeO[x][y] = true
	dfs(board, x+1, y)
	dfs(board, x-1, y)
	dfs(board, x, y-1)
	dfs(board, x, y+1)
}
