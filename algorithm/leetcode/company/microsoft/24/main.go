package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var dummy = &ListNode{}
	dummy.Next = head
	var p1, p2 = dummy, head
	var p *ListNode

	for p2 != nil && p2.Next != nil {
		p2 = p2.Next
		p = p1.Next
		p1.Next = p2
		p.Next = p2.Next
		p2.Next = p
		p1 = p
		p2 = p1.Next
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
