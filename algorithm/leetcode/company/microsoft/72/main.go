package main

func minDistance(word1 string, word2 string) int {
	var m = len(word1)
	var n = len(word2)
	var dp = make([][]int, m+1) // dp[m+1][n+1], word1的前i个字符和word的前j个字符相等的编辑距离
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// dp[i-1][j]:删除word1的一个字符
				// dp[i][j-1]:删除world2的一个字符，或者是插入world1一个字符
				// dp[i-1][j-1]:替换world1或者word2的一个字符
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]

}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b < a && b < c {
		return b
	} else {
		return c
	}
}
