package main

import "math"

var res int

func sum(root *TreeNode) int {
	res = math.MinInt32
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var leftSum = dfs(root.Left)
	var rightSum = dfs(root.Right)
	tmpAns := max2(res, leftSum+root.Val, rightSum+root.Val)
	res = max1(tmpAns, res)
	return tmpAns
}

func max1(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func max2(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= c && b >= a {
		return b
	} else {
		return c
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
