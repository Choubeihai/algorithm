package main

func partition(head *ListNode, x int) *ListNode {
	var dummy = &ListNode{}
	var big = dummy
	var small = dummy
	var p = head
	for p != nil {
		if p.Val >= x {
			p1 := big.Next
			big.Next = p
			p = p.Next
			big.Next.Next = p1
			big = big.Next
		} else {
			p2 := small.Next
			small.Next = p
			p = p.Next
			small.Next.Next = p2
			small = small.Next
			if big.Next.Val < x {
				big = big.Next
			}
		}
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
