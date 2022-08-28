package main

var head *TreeNode // 结果需要返回的头节点
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	head = nil
	dfs(root, nil)
	return head
}

func dfs(root *TreeNode, pre *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left, root)
	if head == nil {
		head = root // 头节点置为最左侧的节点
	}
	if pre != nil {
		pre.Left = nil
		root.Left = pre.Right // 当前节点的左儿子置为父亲的右儿子
		pre.Right = nil
		root.Right = pre // 当前节点的右儿子置为父亲
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
