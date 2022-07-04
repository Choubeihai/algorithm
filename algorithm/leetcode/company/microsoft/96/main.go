package main

import "fmt"

func main() {
	n := 3
	fmt.Println(numTrees(n))
}

/*
   dp五部曲:
   1.状态定义:dp[i]为当有i个节点时,一共可以组成的二叉搜索树数目
   2.状态转移:dp[3]=dp[0]*dp[2]+dp[1]*dp[1]+dp[2]*dp[0]
       可以比喻成前面一项是左子树情况数,后面一项是右子树情况数,相加即可
       即:dp[i]=∑dp[j]*dp[i-1-j],其中j∈[0,i-1]
   3.初始化:dp[0]=1,dp[1]=dp[0]*dp[0]=1
   4.遍历顺序:正序
   4.返回形式:返回dp[n]
*/
func numTrees(n int) int {
	dp := make([]int, n+1) //dp[n]表示i个节点存在的二叉排序树的个数
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			dp[i] = dp[i] + dp[j]*dp[i-1-j]
		}
	}
	return dp[n]
}
