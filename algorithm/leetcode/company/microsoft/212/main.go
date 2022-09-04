package main

var res []string
var visit [][]bool
var dir [][]int
var mm map[string]bool

func findWords(board [][]byte, words []string) []string {
	var m = len(board)
	var n = len(board[0])

	res = nil
	visit = make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	dir = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	mm = make(map[string]bool)
	for i := 0; i < len(words); i++ {
		mm[words[i]] = true
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var b []byte
			visit[i][j] = true
			b = append(b, board[i][j])
			dfs(board, b, i, j)
			visit[i][j] = false
			b = nil
		}
	}
	return res

}

func dfs(board [][]byte, b []byte, x, y int) {
	if len(b) > 10 {
		return
	}
	key := string(b)
	if _, ok := mm[key]; ok {
		delete(mm, key)
		res = append(res, key)
	}
	for i := 0; i < 4; i++ {
		newX := x + dir[i][0]
		newY := y + dir[i][1]
		if newX >= len(board) || newX < 0 || newY >= len(board[0]) || newY < 0 {
			continue
		}
		if visit[newX][newY] {
			continue
		}
		visit[newX][newY] = true
		b = append(b, board[newX][newY])
		dfs(board, b, newX, newY)
		b = b[:len(b)-1]
		visit[newX][newY] = false
	}
}
