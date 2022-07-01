package main

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	if dfs(A, B) {
		return true
	} else if A != nil {
		return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
	} else {
		return false
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
