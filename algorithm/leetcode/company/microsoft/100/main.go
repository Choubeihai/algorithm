package main

/**
判断两棵树是否是完全相同的，必然是递归判断，找好递归出口就好了。
1. 首先如果两者都为空，返回true
2. 只有一方为空，返回false
3. 接下来按照前序 中序 后序遍历都行，这里我以前序为例
4. 先左子树
5. 判断根节点
6. 最后右子树
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return dfs(p, q)
}

func dfs(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil {
		return false
	}
	if q == nil {
		return false
	}
	if !dfs(p.Left, q.Left) {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return dfs(p.Right, q.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
