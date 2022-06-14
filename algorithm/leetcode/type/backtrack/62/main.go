package main

import "fmt"

func main() {
	m := 3
	n := 7
	res := uniquePaths2(m, n)
	fmt.Println(res)
}

// 方案一：动态规划
func uniquePaths(m, n int) int {
	dp := make([][]int, m) // dp[m][n]表示从(0,0)走到(i,j)的路径数量
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1 // 当列只有1行时，那么所有dp[i][1]都为1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1 // 当行只有1行时，那么所有dp[0][i]都为1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 方案二：朴素回溯，会超时
// 朴素的写法： dfs搜索，如果到了右下角，即x==m-1&&y==n-1，计数值就+1；最后输出计数值。但是由于状态重复计算过多，会超时。
var count int
var visit [][]bool // visit也可以不加，因为本身是不会访问到之前访问的节点
var direction [][]int
var mm int
var nn int

func uniquePaths2(m, n int) int {
	count = 0
	visit = make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	direction = [][]int{{1, 0}, {0, 1}}
	mm = m
	nn = n
	dfs(0, 0)
	return count
}

func dfs(x, y int) {
	if x == mm-1 && y == nn-1 {
		count++
	}
	visit[x][y] = true
	for i := 0; i < 2; i++ {
		newX := x + direction[i][0]
		newY := y + direction[i][1]
		if newX >= 0 && newX < mm && newY >= 0 && newY < nn && !visit[newX][newY] {
			dfs(newX, newY)
		}
	}
	visit[x][y] = false
}

// 回溯
