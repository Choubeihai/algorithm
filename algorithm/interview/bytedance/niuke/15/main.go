package main

import (
	"fmt"
)

/**
给两个数组，长度相等，一个只能pop front,一个可以pop front和 pop back,
两个数组每次各自pop出一个数a和b，sum+=a*b，直至两数组取完，求sum最大值
*/

func main() {
	num1 := []int{1, 2, 3, 4}
	num2 := []int{4, 3, 2, 1}
	fmt.Println(getMaxSum(num1, num2))
}

func getMaxSum(num1 []int, num2 []int) int {
	var n = len(num1)
	var dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = num1[i] * num2[i]
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			m1 := dp[i+1][j] + num1[i]*num2[i]
			m2 := dp[i][j-1] + num1[i]*num2[j]
			dp[i][j] = max(m1, m2)
		}
	}

	return dp[0][n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
