package main

func main() {

}

func maxHeight() {

}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := dfs(root.Left)
	rightHeight := dfs(root.Right)
	return 1 + max(leftHeight, rightHeight)
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
