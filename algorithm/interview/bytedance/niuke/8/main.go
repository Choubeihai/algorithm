package main

/**
剑指offer 26
一面上来就是两道算法题，第一道是给两颗二叉树，判断B是不是A的子树。
*/

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	if A == nil {
		return false
	}
	if dfs(A, B) {
		return true
	} else {
		return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
	}
}

func dfs(root1 *TreeNode, root2 *TreeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}

	return dfs(root1.Left, root2.Left) && dfs(root1.Right, root2.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
