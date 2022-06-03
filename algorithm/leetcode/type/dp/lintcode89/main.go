package main

import "fmt"

/**
给定 n 个不同的正整数，整数 k(k≤n)以及一个目标数字 target。
在这 n 个数里面找出 k 个数，使得这 k 个数的和等于目标数字，求问有多少种方案？
*/

func main() {
	var a = []int{1, 2, 3, 4}
	k := 2
	target := 5
	fmt.Println(KSum(a, k, target))
}

func KSum(a []int, k int, target int) int {
	if target < 0 {
		return 0
	}
	n := len(a)
	var dp [][][]int          // f(x,y,z): 从x个数中找出y个数，和为z的方案数
	dp = make([][][]int, n+1) // dp[n+1][k+1][target+1]
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]int, target+1)
		}
	}
	for i := 0; i <= n; i++ {
		dp[i][0][0] = 1
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= k; j++ {
			for l := 0; l <= target; l++ {
				if j == 0 && l == 0 {
					dp[i][j][l] = 1
				} else if i != 0 && j != 0 && l != 0 {
					dp[i][j][l] = dp[i-1][j][l]
					if l >= a[i-1] {
						dp[i][j][l] = dp[i][j][l] + dp[i-1][j-1][l-a[i-1]]
					}
				}
			}
		}
	}
	return dp[n][k][target]
}
