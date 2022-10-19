package main

var dp [][]bool
var res [][]string

func partition(s string) [][]string {
	n := len(s)
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
	dfs(s, 0, []string{})
	return res
}

func dfs(s string, startIdx int, b []string) {
	if startIdx == len(s) {
		c := make([]string, len(b))
		copy(c, b)
		res = append(res, c)
		return
	}
	for i := startIdx; i < len(s); i++ {
		if dp[startIdx][i] == false {
			continue
		} else {
			b = append(b, s[startIdx:i+1])
			dfs(s, i+1, b)
			b = b[:len(b)-1]
		}
	}
}
