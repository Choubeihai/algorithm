package main

func maxProfit(prices []int) int {
	var res int
	var minPrice = prices[0]
	for i := 1; i < len(prices); i++ {
		res = max(res, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
