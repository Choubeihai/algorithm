package main

var tag [][]bool

func solve(board [][]byte) {
	var m = len(board)
	var n = len(board[0])
	tag = make([][]bool, m)
	for i := 0; i < m; i++ {
		tag[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || i == m-1 {
				dfs(board, i, j)
			} else if j == 0 || j == n-1 {
				dfs(board, i, j)
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if tag[i][j] == false && board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(grid [][]byte, x, y int) {
	if x >= len(grid) || x < 0 || y >= len(grid[0]) || y < 0 {
		return
	}
	if grid[x][y] == 'X' {
		return
	}
	if tag[x][y] == true {
		return
	}
	tag[x][y] = true
	dfs(grid, x, y+1)
	dfs(grid, x+1, y)
	dfs(grid, x, y-1)
	dfs(grid, x-1, y)
}
