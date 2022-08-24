package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var m map[*TreeNode]int
var n map[*TreeNode]int

func rob(root *TreeNode) int {
	m = make(map[*TreeNode]int) // 包含root
	n = make(map[*TreeNode]int) // 不包含root
	dfs(root)
	return max(m[root], n[root])
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	dfs(root.Right)
	m[root] = root.Val + n[root.Left] + n[root.Right]
	n[root] = max(m[root.Left], n[root.Left]) + max(m[root.Right], n[root.Right])
}

func max(a, b int) int {
	if a >= b {
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
