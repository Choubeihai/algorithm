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
	/***
	为什么用dp(0， n-1)表示结果？
	 dp[0][n-1]表示的是完整的初始数组，游戏刚开始拿的一定是先手玩家呀。
	能在[0,n-1]范围内拿到分差大于0，就是先手赢得意思。 如果按照你说的后手视角，那数组就不是从0到n-1了，因为数组已经被拿过了。
	*/
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
