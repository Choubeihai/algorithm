package main

func rob(nums []int) int {
	var n = len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	var dp = make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n-1; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	var max1 = dp[n-2]

	dp[1] = nums[1]
	dp[2] = max(nums[1], nums[2])
	for i := 3; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return max(max1, dp[n-1])

}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
