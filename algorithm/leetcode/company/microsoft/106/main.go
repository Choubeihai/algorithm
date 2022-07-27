package main

func sortedListToBST(head *ListNode) *TreeNode {

	return build(head)
}

func build(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var preSlow *ListNode
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	root := &TreeNode{}
	root.Val = slow.Val
	if preSlow != nil {
		preSlow.Next = nil
		root.Left = build(head)
	}

	root.Right = build(slow.Next)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
