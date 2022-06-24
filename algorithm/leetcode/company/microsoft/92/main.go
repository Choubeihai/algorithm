package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	var dummy = &ListNode{}
	dummy.Next = head
	var p = dummy
	var p1Tail *ListNode
	var p2 *ListNode
	var p3 *ListNode
	var index = 0
	for p != nil {
		if index == left-1 {
			p1Tail = p
			p2 = p.Next
			p.Next = nil
			p = p2
			index++
		} else if index == right {
			p3 = p.Next
			p.Next = nil
			index++
		} else {
			p = p.Next
			index++
		}
	}
	p2Head, p2Tail := reverse(p2)
	p1Tail.Next = p2Head
	p2Tail.Next = p3
	return dummy.Next
}

func reverse(head *ListNode) (newHead, tail *ListNode) {
	var p = head
	var pre *ListNode
	for p != nil {
		p1 := p.Next
		p.Next = pre
		pre = p
		p = p1
	}
	return pre, head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
