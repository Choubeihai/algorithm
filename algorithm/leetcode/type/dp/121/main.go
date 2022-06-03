package main

func main() {

}

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	var n = len(prices)
	var dp = make([]int, n+1) // dp表示当然能赚取的最大利润, i表示第几天
	dp[0] = 0
	var minPrice = prices[0]
	for i := 1; i <= n; i++ {
		if minPrice > prices[i-1] {
			minPrice = prices[i-1]
			dp[i] = dp[i-1]
		} else {
			dp[i] = max(dp[i-1], prices[i-1]-minPrice)
		}
	}
	var maxProfit = dp[0]
	for i := 1; i < len(dp); i++ {
		if maxProfit < dp[i] {
			maxProfit = dp[i]
		}
	}
	return maxProfit
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
