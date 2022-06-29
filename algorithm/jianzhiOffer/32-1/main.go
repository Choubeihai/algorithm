package main

func levelOrder(root *TreeNode) []int {
	var queue = []*TreeNode{root}
	var res []int
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			res = append(res, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
