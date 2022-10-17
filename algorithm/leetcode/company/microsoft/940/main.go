package main

func distinctSubseqII(s string) int {
	dp := make([]int, 26)
	for i := 0; i < len(s); i++ {
		index := int(s[i] - 'a')
		dp[index] = calculate(dp) + 1

	}
	res := calculate(dp)
	return res
}

func calculate(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = (sum + arr[i]) % (1e9 + 7)
	}
	return sum
}
