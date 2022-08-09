package main

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var p1 = head
	var p2 = head.Next
	var evenHead = p2
	for p2 != nil && p2.Next != nil {
		p1.Next = p2.Next
		p1 = p1.Next
		p2.Next = p1.Next
		p2 = p2.Next
	}
	p1.Next = evenHead
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
