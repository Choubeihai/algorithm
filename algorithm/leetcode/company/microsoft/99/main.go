package main

var pre *TreeNode
var big *TreeNode
var small *TreeNode

func recoverTree(root *TreeNode) {
	pre = nil
	big = nil
	small = nil
	dfs(root)
	small.Val, big.Val = big.Val, small.Val
}

// 中序遍历
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if pre != nil && pre.Val > root.Val {
		if big == nil {
			big = pre
			small = root
		} else {
			small = root
		}
	}
	pre = root
	dfs(root.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
