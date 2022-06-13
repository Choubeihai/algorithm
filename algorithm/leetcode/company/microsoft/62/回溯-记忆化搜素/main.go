package main

/**
朴素的写法： dfs搜索，如果到了右下角，即x==m-1&&y==n-1，计数值就+1；最后输出计数值。但是由于状态重复计算过多，会超时。
记忆化：利用一个二维数组mem存当前节点可以达到右下角的路径数，这样每次变化newX, newY想要去搜索一个新的节点的时候
       我们可以利用曾经计算过的结果，节约了很多计算时间。

*/

var visit [][]bool
var mem [][]int // 当前节点可以达到右下角的路径数
var direction [][]int
var mm int
var nn int

func uniquePaths(m, n int) int {
	visit = make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}

	mem = make([][]int, m)
	for i := 0; i < m; i++ {
		mem[i] = make([]int, n)
	}

	direction = [][]int{{1, 0}, {0, 1}}
	mm = m
	nn = n
	res := dfs(0, 0)
	return res
}

func dfs(x, y int) int {
	if x == mm-1 && y == nn-1 {
		return 1
	}
	visit[x][y] = true
	for i := 0; i < 2; i++ {
		newX := x + direction[i][0]
		newY := y + direction[i][1]
		if newX >= 0 && newX < mm && newY >= 0 && newY < nn && !visit[newX][newY] {
			if mem[newX][newY] != 0 {
				mem[x][y] = mem[x][y] + mem[newX][newY]
			} else {
				mem[x][y] = mem[x][y] + dfs(newX, newY)
			}
		}
	}
	visit[x][y] = false
	return mem[x][y]
}
