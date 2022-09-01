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

// 方法二
func swapPairs2(head *ListNode) *ListNode {
	var dummy = &ListNode{}
	dummy.Next = head
	var p = dummy
	for p.Next != nil && p.Next.Next != nil {
		p1 := p.Next.Next.Next
		p.Next.Next.Next = nil

		p2 := p.Next
		p.Next = p.Next.Next
		p.Next.Next = p2
		p2.Next = p1
		p = p2
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
