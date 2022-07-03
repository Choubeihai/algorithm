package main

// 和lc 10题一样
func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		dp[i][0] = false
	}
	for i := 1; i <= n; i++ {
		if p[i-1] != '*' {
			dp[0][i] = false
		} else {
			dp[0][i] = dp[0][i-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// dp[i][j-2]表示*代表0个字符，所以取dp[i][j-2]
				// dp[i-1][j]表示*代表1个或者1个以上字符，所以将s的字符删除一个，仍然可以匹配，取dp[i-1][j]
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i-1][j] || dp[i][j-2]
				} else {
					// // dp[i][j-2]表示*代表0个字符，所以取dp[i][j-2]
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[m][n]
}
