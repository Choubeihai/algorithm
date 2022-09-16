package main

import "fmt"

func main() {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}

// 动态规划
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	var m = len(obstacleGrid)
	var n = len(obstacleGrid[0])
	var dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	flag := false
	for i := 0; i < m; i++ {
		if flag {
			dp[i][0] = 0
		} else if obstacleGrid[i][0] == 1 {
			flag = true
			dp[i][0] = 0
		} else {
			dp[i][0] = 1
		}
	}
	flag = false
	for i := 0; i < n; i++ {
		if flag {
			dp[0][i] = 0
		} else if obstacleGrid[0][i] == 1 {
			flag = true
			dp[0][i] = 0
		} else {
			dp[0][i] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 记忆化搜索
var memo [][]int // 到达x,y点时有几条方案。
var dir [][]int
var m int
var n int

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m = len(obstacleGrid)
	n = len(obstacleGrid[0])
	dir = [][]int{{1, 0}, {0, 1}}
	memo = make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}
	res := dfs(obstacleGrid, 0, 0)
	return res
}

func dfs(obstacleGrid [][]int, x, y int) int {
	if x >= m || x < 0 || y >= n || y < 0 {
		return 0
	}
	if obstacleGrid[x][y] == 1 {
		return 0
	}
	if x == len(obstacleGrid)-1 && y == len(obstacleGrid[0])-1 {
		return 1
	}

	memo[x][y] = 0
	for i := 0; i < 2; i++ {
		newX := x + dir[i][0]
		newY := y + dir[i][1]
		if newX >= 0 && newY >= 0 && newX <= m-1 && newY <= n-1 {
			if memo[newX][newY] != -1 {
				memo[x][y] += memo[newX][newY]
			} else {
				memo[newX][newY] = dfs(obstacleGrid, newX, newY)
				memo[x][y] += memo[newX][newY]
			}
		}
	}
	return memo[x][y]
}
