package main

import "fmt"

func main() {
	s1 := "db"
	s2 := "b"
	s3 := "cbb"
	fmt.Println(isInterleave(s1, s2, s3))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if (n + m) != t {
		return false
	}

	dp := make([][]bool, m+1) // dp[m][n] 表示s1的前m个字符和s2的前n个字符是否可以组成s3的钱n+m个字符
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	for i := 1; i <= m; i++ {
		if s3[i-1] == s1[i-1] {
			dp[i][0] = dp[i-1][0]
		} else {
			dp[i][0] = false
		}
	}
	for i := 1; i <= n; i++ {
		if s3[i-1] == s2[i-1] {
			dp[0][i] = dp[0][i-1]
		} else {
			dp[0][i] = false
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[m][n]
}
