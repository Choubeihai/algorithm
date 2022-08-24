package main

var encoding string

func numDistinctIslands(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	var mm = make(map[string]int)
	encoding = ""
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				encoding = ""
				dfs(grid, i, j, "s") // 打上标签
				mm[encoding] = 1
			}
		}
	}
	return len(mm)
}

func dfs(grid [][]int, x, y int, directionTag string) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return
	}
	if grid[x][y] == 0 { // 碰到水了
		return
	}

	grid[x][y] = 0
	encoding = encoding + directionTag
	dfs(grid, x, y+1, "r")    // 向右
	dfs(grid, x+1, y, "d")    // 向下
	dfs(grid, x, y-1, "l")    // 向左
	dfs(grid, x-1, y, "u")    // 向上
	encoding = encoding + "e" // 结束
}
