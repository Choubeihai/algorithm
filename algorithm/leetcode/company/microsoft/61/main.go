package main

func rotateRight(head *ListNode, k int) *ListNode {
	var n = 0
	var dummy = &ListNode{}
	dummy.Next = head
	var p = head
	for p != nil {
		p = p.Next
		n++
	}
	if n == 0 {
		return nil
	}
	k = n - k%n

	var p1 = dummy
	var p2 = dummy
	var cnt = 0

	for cnt != k {
		cnt++
		p1 = p1.Next
	}

	for p1 != nil && p1.Next != nil {
		p = p1.Next
		p1.Next = p.Next

		pp := p2.Next
		p2.Next = p
		p.Next = pp
		p2 = p2.Next
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
