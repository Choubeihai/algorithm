package main

func reverseList(head *ListNode) *ListNode {
	return reverse(head)
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		p := head.Next
		head.Next = pre
		pre = head
		head = p
	}
	return pre
}

type ListNode struct {
	Val  int
	Next *ListNode
}
