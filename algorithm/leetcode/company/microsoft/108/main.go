package main

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	return buid(nums, 0, n-1)
}

func buid(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	var root = &TreeNode{}
	root.Val = nums[mid]
	root.Left = buid(nums, left, mid-1)
	root.Right = buid(nums, mid+1, right)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
