package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy = &ListNode{}
	p := dummy
	bit := 0 // 进位
	for l1 != nil && l2 != nil {
		value := l1.Val + l2.Val + bit
		bit = value / 10
		value = value % 10
		p.Next = &ListNode{Val: value}
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		value := l1.Val + bit
		bit = value / 10
		value = value % 10
		p.Next = &ListNode{Val: value}
		p = p.Next
		l1 = l1.Next
	}

	for l2 != nil {
		value := l2.Val + bit
		bit = value / 10
		value = value % 10
		p.Next = &ListNode{Val: value}
		p = p.Next
		l1 = l1.Next
	}

	if bit != 0 {
		p.Next = &ListNode{Val: bit}
	}

	return dummy.Next
}
