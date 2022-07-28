package main

func minDepth(root *TreeNode) int {
	return dfs(root)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := dfs(root.Left)
	rightDepth := dfs(root.Right)
	if root.Left != nil && root.Right != nil {
		return 1 + min(leftDepth, rightDepth)
	} else if root.Right == nil {
		return 1 + leftDepth
	} else {
		return 1 + rightDepth
	}

}
func min(a, b int) int {
	if a < b {
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
