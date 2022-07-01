package main

var myTarget int
var res [][]int

func pathSum(root *TreeNode, target int) [][]int {
	myTarget = target
	res = nil
	dfs(root, []int{}, target)
	return res
}

func dfs(root *TreeNode, b []int, sum int) {
	if root == nil {
		return
	}
	sum -= root.Val
	b = append(b, root.Val)
	if root.Left == nil && root.Right == nil && sum == 0 {
		c := make([]int, len(b))
		copy(c, b)
		res = append(res, c)
	}
	dfs(root.Left, b, sum)
	dfs(root.Right, b, sum)
	sum += root.Val
	b = b[:len(b)-1]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
