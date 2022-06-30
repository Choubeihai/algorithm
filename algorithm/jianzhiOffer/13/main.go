package main

var direction [][]int
var visit [][]bool
var res int
var kk int
var mm int
var nn int

func movingCount(m int, n int, k int) int {
	direction = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	res = 0
	mm = m
	nn = n
	kk = k
	visit = make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	dfs(0, 0)
	return res
}

func dfs(x, y int) {
	if x >= mm || y >= nn || x < 0 || y < 0 {
		return
	}
	if visit[x][y] {
		return
	}
	xx := x
	yy := y
	sum := 0
	for xx != 0 {
		sum += xx % 10
		xx /= 10
	}
	for yy != 0 {
		sum += yy % 10
		yy /= 10
	}
	if sum > kk {
		return
	} else {
		visit[x][y] = true
		res++
	}
	for i := 0; i < 4; i++ {
		newX := x + direction[i][0]
		newY := y + direction[i][1]
		dfs(newX, newY)
	}
}
