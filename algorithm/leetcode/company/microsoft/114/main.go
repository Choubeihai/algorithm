package main

func flatten(root *TreeNode) {
	dfs(root)
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		right := root.Right
		cur := root.Left
		for cur.Right != nil {
			cur = cur.Right
		}
		cur.Right = right
		root.Right = root.Left
		root.Left = nil
	}
	dfs(root.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
