package main

var mm map[[2]int]int
var visit [][]bool
var dir [][]int
var res int

func uniquePathsIII(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	mm = make(map[[2]int]int)
	visit = make([][]bool, m)
	dir = [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}
	res = 0
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	x := 0
	y := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				x = i
				y = j
			}
			if grid[i][j] == 0 {
				mm[[2]int{i, j}] = 1
			}
		}
	}
	visit[x][y] = true
	dfs(grid, x, y)
	return res
}

func dfs(grid [][]int, x, y int) {
	if x >= len(grid) || x < 0 || y >= len(grid[0]) || y < 0 {
		return
	}
	if grid[x][y] == 2 {
		if len(mm) == 0 {
			res++
		}
		return
	}
	if grid[x][y] == -1 {
		return
	}
	for i := 0; i < 4; i++ {
		xx := x + dir[i][0]
		yy := y + dir[i][1]
		if xx >= 0 && xx < len(grid) && yy >= 0 && yy < len(grid[0]) && !visit[xx][yy] {
			visit[xx][yy] = true
			if grid[xx][yy] == 0 {
				delete(mm, [2]int{xx, yy})
			}
			dfs(grid, xx, yy)
			visit[xx][yy] = false
			if grid[xx][yy] == 0 {
				mm[[2]int{xx, yy}] = 1
			}
		}
	}
}
