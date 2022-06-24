package main

import "fmt"

func main() {
	n := 3
	fmt.Println(numTrees(n))
}

func numTrees(n int) int {
	dp := make([]int, n+1) //dp[n]表示i个节点存在的二叉排序树的个数
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
