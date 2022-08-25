package main

func insertionSortList(head *ListNode) *ListNode {
	return sort(head)
}

func sort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var fast = head
	var slow = head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	p := slow.Next
	slow.Next = nil
	head = sort(head)
	head2 := sort(p)
	return merge(head, head2)
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy = &ListNode{}
	var p = dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
			p = p.Next
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
