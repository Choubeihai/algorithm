package main

import "fmt"

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	var dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < n+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < m+1; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[n][m]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
