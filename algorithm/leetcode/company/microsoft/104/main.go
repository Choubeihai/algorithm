package main

func maxDepth(root *TreeNode) int {
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(dfs(root.Left), dfs(root.Right))
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
