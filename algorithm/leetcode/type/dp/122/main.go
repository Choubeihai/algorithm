package main

import "fmt"

func main() {
	var prices = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var n = len(prices)
	// dp[n+1][2]:表示目前手上最大的盈利，i表示第几天，j表示是否持有股票
	var dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[1][0] = 0
	dp[1][1] = -prices[0]
	for i := 2; i <= n; i++ {
		dp[i][0] = max(dp[i-1][1]+prices[i-1], dp[i-1][0])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i-1])
	}
	return dp[n][0]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
