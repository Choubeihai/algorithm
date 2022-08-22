package main

// 和486题很像
func stoneGame(piles []int) bool {
	var n = len(piles)
	var dp = make([][]int, n) // f(x, y)：在区间[x,y]中，当前玩家与另一个玩家的差值
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = piles[i] // 当i=j时，只剩一个数字，当前玩家只能拿取这个数字
	}

	// 保证右边界j比左边界i大
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			m1 := piles[i] - dp[i+1][j]
			m2 := piles[j] - dp[i][j-1]
			dp[i][j] = max(m1, m2)
		}
	}
	if dp[0][n-1] > 0 {
		return true
	} else {
		return false
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
