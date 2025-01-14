package main

func rob(nums []int) int {
	var n = len(nums)
	if n == 1 {
		return nums[0]
	}
	var dp = make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
