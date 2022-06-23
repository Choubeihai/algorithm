package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var dummy = &ListNode{Val: 101}
	var p = dummy
	dummy.Next = head
	for p != nil && p.Next != nil {
		if p.Next.Next != nil && p.Next.Next.Val == p.Next.Val {
			value := p.Next.Val
			p1 := p.Next.Next.Next
			for p1 != nil && p1.Val == value {
				p1 = p1.Next
			}
			p.Next = p1
		} else {
			p = p.Next
		}
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
