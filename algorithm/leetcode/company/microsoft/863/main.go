package main

var parents map[int]*TreeNode
var res []int

func distanceK(root, target *TreeNode, k int) (ans []int) {
	res = nil
	parents = make(map[int]*TreeNode)
	findParents(root)
	dfs(target, nil, 0, k)
	return res
}

// 从 root 出发 DFS，记录每个结点的父结点
func findParents(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		parents[root.Left.Val] = root
	}
	if root.Right != nil {
		parents[root.Right.Val] = root
	}
	findParents(root.Left)
	findParents(root.Right)
}

// 从 target 出发 DFS，寻找所有深度为 k 的结点
func dfs(root *TreeNode, from *TreeNode, distance int, k int) {
	if root == nil {
		return
	}
	if distance == k {
		res = append(res, root.Val)
		return
	}
	if root.Left != from {
		dfs(root.Left, root, distance+1, k)
	}
	if root.Right != from {
		dfs(root.Right, root, distance+1, k)
	}
	if parents[root.Val] != from {
		dfs(parents[root.Val], root, distance+1, k)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
