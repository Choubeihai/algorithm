package main

var res int
var area int

func maxAreaOfIsland(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	res = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			area = 0
			if grid[i][j] == 1 {
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]int, x, y int) {
	if x >= len(grid) || x < 0 || y >= len(grid[0]) || y < 0 {
		return
	}
	if grid[x][y] == 0 {
		return
	}
	grid[x][y] = 0
	area++
	dfs(grid, x, y+1)
	dfs(grid, x+1, y)
	dfs(grid, x, y-1)
	dfs(grid, x-1, y)
	if res < area {
		res = area
	}
}
