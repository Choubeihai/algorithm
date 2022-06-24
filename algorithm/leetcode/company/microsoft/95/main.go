package main

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return build(1, n)
}

func build(left, right int) []*TreeNode {
	if left > right {
		return []*TreeNode{nil}
	}
	var res []*TreeNode
	for i := left; i <= right; i++ {
		leftTreeList := build(left, i-1)
		rightTreeList := build(i+1, right)
		for j := 0; j < len(leftTreeList); j++ {
			for k := 0; k < len(rightTreeList); k++ {
				root := &TreeNode{Val: i}
				root.Left = leftTreeList[j]
				root.Right = rightTreeList[k]
				res = append(res, root)
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
