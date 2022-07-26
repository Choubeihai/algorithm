package main

func dfs(root *TreeNode, target int) (int, int) {
	var level = 0
	var count = 0
	var q []*TreeNode
	q = append(q, root)
	for len(q) != 0 {
		level++
		length := len(q)
		for i := 0; i < length; i++ {
			count++
			if q[i].Val == target {
				return level, count
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[length:]
	}
	return -1, -1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
