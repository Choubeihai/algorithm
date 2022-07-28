package main

var res [][]int
var target int

func pathSum(root *TreeNode, targetSum int) [][]int {
	res = nil
	target = targetSum
	dfs(root, []int{}, 0)
	return res
}

func dfs(root *TreeNode, path []int, sum int) {
	if root == nil {
		return
	}
	sum += root.Val
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && sum == target {
		c := make([]int, len(path))
		copy(c, path)
		res = append(res, c)
	}
	dfs(root.Left, path, sum)
	dfs(root.Right, path, sum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
