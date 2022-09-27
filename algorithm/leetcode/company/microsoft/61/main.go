package main

func rotateRight(head *ListNode, k int) *ListNode {
	p := head
	n := 0
	for p != nil {
		p = p.Next
		n++
	}
	if n == 0 {
		return head
	}
	k = k % n
	dummy := &ListNode{}
	dummy.Next = head
	p1 := dummy
	p2 := dummy
	p = dummy
	for k > 0 {
		k--
		p1 = p1.Next
	}
	for p1 != nil && p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	for p2 != nil && p2.Next != nil {
		p3 := p.Next
		p.Next = p2.Next
		p2.Next = p2.Next.Next
		p.Next.Next = p3
		p = p.Next
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
