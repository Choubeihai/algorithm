package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	n := getLevels(root)
	var res [][]int = make([][]int, n)
	var q []*TreeNode
	q = append(q, root)
	for len(q) != 0 {
		length := len(q)
		var tmpRes []int
		for i := 0; i < length; i++ {
			tmpRes = append(tmpRes, q[i].Val)
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[length:]
		n--
		res[n] = tmpRes
	}
	return res
}

func getLevels(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getLevels(root.Left), getLevels(root.Right))
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
