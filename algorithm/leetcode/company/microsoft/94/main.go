package main

var res []int

func inorderTraversal(root *TreeNode) []int {
	res = nil
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	res = append(res, root.Val)
	dfs(root.Right)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
