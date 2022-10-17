package main

func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n)

	for i := 0; i < len(dislikes); i++ {
		x := dislikes[i][0] - 1
		y := dislikes[i][1] - 1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	colors := make([]int, n) // color = 0 : 表示尚未访问(染色)

	for i := 0; i < len(colors); i++ {
		if colors[i] == 0 && !dfs(i, 1, colors, g) {
			return false
		}
	}
	return true
}

func dfs(node int, c int, colors []int, g [][]int) bool {
	colors[node] = c
	for i := 0; i < len(g[node]); i++ {
		if colors[g[node][i]] == c {
			return false
		}
		if colors[g[node][i]] == 0 && !dfs(g[node][i], 2^c, colors, g) {
			return false
		}
	}
	return true
}
