package main

import "math"

var memo [][]int
var m int
var n int
var myGrid [][]int
var direction = [][]int{{1, 0}, {0, 1}}

// 记忆化搜索
func minPathSum2(grid [][]int) int {
	myGrid = grid
	m = len(grid)
	n = len(grid[0])
	memo = make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		// 初始化memo
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	res := dfs(0, 0)
	return res

}

func dfs(x, y int) int {
	if memo[x][y] != -1 {
		return memo[x][y]
	}
	if x == m-1 && y == n-1 {
		return myGrid[x][y]
	}
	path := math.MaxInt32
	for i := 0; i < 2; i++ {
		newX := x + direction[i][0]
		newY := y + direction[i][1]
		if newX >= 0 && newX < m && newY >= 0 && newY < n {
			path = min(path, dfs(newX, newY))
		}
	}
	memo[x][y] = path + myGrid[x][y]
	return memo[x][y]
}

// 动态规划
func minPathSum(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
