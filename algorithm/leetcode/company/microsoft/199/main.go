package main

// 方法二：bfs
func rightSideView1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var q []*TreeNode
	q = append(q, root)
	for len(q) != 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			node := q[i]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			if i == length-1 {
				res = append(res, q[i].Val)
			}
		}
		q = q[length:]
	}
	return res
}

// 方法一：dfs
var res []int

func rightSideView(root *TreeNode) []int {
	res = nil
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if len(res) == depth {
		res = append(res, root.Val)
	}
	dfs(root.Right, depth+1)
	dfs(root.Left, depth+1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
