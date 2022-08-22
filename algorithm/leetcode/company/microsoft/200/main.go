package main

var res int
var directions [][]int

func numIslands(grid [][]byte) int {
	var m = len(grid)
	var n = len(grid[0])
	res = 0
	directions = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, x, y int) {
	if x >= len(grid) || x < 0 || y >= len(grid[0]) || y < 0 || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	for i := 0; i < 4; i++ {
		newX := directions[i][0] + x
		newY := directions[i][1] + y
		dfs(grid, newX, newY)
	}
}
