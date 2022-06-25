package main

var pre *TreeNode

func isValidBST(root *TreeNode) bool {
	pre = nil
	return dfs(root)
}

func dfs(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !dfs(root.Left) {
		return false
	}

	if pre != nil {
		if pre.Val >= root.Val {
			return false
		}
	}

	pre = root
	return dfs(root.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
