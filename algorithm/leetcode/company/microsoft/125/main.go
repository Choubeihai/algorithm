package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var sum int

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum = 0
	dfs(root, 0)
	return sum
}

func dfs(root *TreeNode, v int) {
	if root == nil {
		return
	}
	v = v*10 + root.Val
	if root.Left == nil && root.Right == nil {
		sum += v
	}
	dfs(root.Left, v)
	dfs(root.Right, v)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
