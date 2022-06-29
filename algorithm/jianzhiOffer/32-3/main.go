package main

func levelOrder(root *TreeNode) [][]int {
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
		if len(res)%2 != 0 {
			left := 0
			right := len(tmpRes) - 1
			for left < right {
				tmpRes[left], tmpRes[right] = tmpRes[right], tmpRes[left]
				left++
				right--
			}
		}
		res = append(res, tmpRes)
		q = q[length:]
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
