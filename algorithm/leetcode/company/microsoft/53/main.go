package main

// 动态规划
func maxSubArray(nums []int) int {
	var n = len(nums)
	var dp = make([]int, n) // dp[i]：表示以index i为结尾的最大连续和(即这个和中必须包括index i对应的值)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}
	var res = dp[0]
	for i := 1; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
