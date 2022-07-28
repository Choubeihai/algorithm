package main

var target int

func hasPathSum(root *TreeNode, targetSum int) bool {
	target = targetSum
	return dfs(root, 0)
}

func dfs(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum += root.Val
	if sum == target && root.Left == nil && root.Right == nil {
		return true
	}
	return dfs(root.Left, sum) || dfs(root.Right, sum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
