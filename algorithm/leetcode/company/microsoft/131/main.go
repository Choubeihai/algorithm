package main

var dp [][]bool
var ss string
var res [][]string

func partition(s string) [][]string {
	n := len(s)
	ss = s
	res = nil
	dp = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for j := 0; j < n; j++ { // 右边界
		for i := 0; i <= j; i++ { // 左边界
			if j-i <= 2 {
				if s[i] == s[j] {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			} else {
				if s[i] != s[j] {
					dp[i][j] = false
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
		}
	}
	dfs(0, []string{})
	return res
}

func dfs(index int, path []string) {
	if index == len(ss) {
		var ans = make([]string, len(path))
		copy(ans, path)
		res = append(res, ans)
		return
	}
	for i := index; i < len(ss); i++ {
		if !dp[index][i] {
			continue
		} else {
			subS := ss[index : i+1]
			path = append(path, subS)
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}
}
