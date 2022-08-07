package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return dfs(root, p, q)
}

func dfs(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := dfs(root.Left, p, q)
	right := dfs(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else {
		return right
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
