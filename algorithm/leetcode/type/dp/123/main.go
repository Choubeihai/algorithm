package main

import "fmt"

func main() {
	var prices = []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	var n = len(prices)
	var m = 2
	// dp[n+1][m+1][2] 表示获取的最大利润，i表示第几天，j表示第几次购买股票，k表示是否持有股票
	var dp = make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = make([]int, 2)

		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if i == 1 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i-1]
			} else {
				dp[i][j][0] = max(dp[i-1][j][1]+prices[i-1], dp[i-1][j][0])
				dp[i][j][1] = max(dp[i-1][j-1][0]-prices[i-1], dp[i-1][j][1])
			}
		}
	}
	return dp[n][m][0]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
