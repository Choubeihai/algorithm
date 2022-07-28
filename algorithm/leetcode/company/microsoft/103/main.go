package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var q []*TreeNode
	q = append(q, root)
	for len(q) != 0 {
		length := len(q)
		var tmpRes []int
		for i := 0; i < length; i++ {
			node := q[i]
			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[length:]
		if len(res)%2 != 0 {
			for i, j := 0, len(tmpRes)-1; i < j; i, j = i+1, j-1 {
				tmpRes[i], tmpRes[j] = tmpRes[j], tmpRes[i]
			}
		}
		res = append(res, tmpRes)

	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
