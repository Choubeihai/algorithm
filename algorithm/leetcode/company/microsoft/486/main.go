package main

func PredictTheWinner(nums []int) bool {
	var n = len(nums)
	var dp = make([][]int, n) // f(x, y)：在区间[x,y]中，当前玩家与另一个玩家的差值
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = nums[i]
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			m1 := nums[i] - dp[i+1][j]
			m2 := nums[j] - dp[i][j-1]
			dp[i][j] = max(m1, m2)
		}
	}
	if dp[0][n-1] >= 0 {
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
