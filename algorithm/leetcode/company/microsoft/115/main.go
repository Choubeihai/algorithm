package main

func numDistinct(s string, t string) int {
	var m = len(s)
	var n = len(t)
	if n > m {
		return 0
	}
	var dp = make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := 1; i < n+1; i++ {
		dp[0][i] = 0
	}
	for i := 1; i < m+1; i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}
