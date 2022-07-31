package main

import (
	"math"
)

func minimumTotal(triangle [][]int) int {
	var n = len(triangle)
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if i-1 >= 0 {
				if j <= i-1 {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = math.MaxInt32 - 100
				}
				if j-1 >= 0 {
					dp[i][j] = min(dp[i][j], dp[i-1][j-1])
				}
				dp[i][j] += triangle[i][j]
			} else {
				dp[i][j] = triangle[i][j]
			}
		}
	}
	var res = math.MaxInt32
	for i := 0; i < n; i++ {
		if res > dp[n-1][i] {
			res = dp[n-1][i]
		}

	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
