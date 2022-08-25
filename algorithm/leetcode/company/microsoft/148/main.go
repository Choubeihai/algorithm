package main

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	head2 := slow.Next
	slow.Next = nil
	return merge(sortList(head), sortList(head2))
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy = &ListNode{}
	var p = dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			p = p.Next
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
			p = p.Next
		}
	}
	for l1 != nil {
		p.Next = l1
		l1 = l1.Next
		p = p.Next
	}

	for l2 != nil {
		p.Next = l2
		l2 = l2.Next
		p = p.Next
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
