package main

func sortedListToBST(head *ListNode) *TreeNode {
	return build(head)
}

func build(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	var preSlow *ListNode
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if preSlow != nil {
		preSlow.Next = nil
	}
	root := &TreeNode{}
	root.Val = slow.Val
	root.Left = build(head)
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
