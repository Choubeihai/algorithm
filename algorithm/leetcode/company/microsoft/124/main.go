package main

import "math"

var maxSum int

func maxPathSum(root *TreeNode) int {
	maxSum = math.MinInt32
	dfs(root)
	return maxSum
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftSum := max(dfs(root.Left), 0)
	rightSum := max(dfs(root.Right), 0)
	sum := root.Val + leftSum + rightSum
	maxSum = max(maxSum, sum)
	return max(leftSum, rightSum) + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
