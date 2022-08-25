package main

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var fast = head
	var slow = head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	p := slow.Next
	slow.Next = nil
	head2 := reverse(p)
	merge(head, head2)
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy = &ListNode{}
	var p = dummy
	var count = 0
	for l1 != nil && l2 != nil {
		if count%2 == 0 {
			count++
			p.Next = l1
			p = p.Next
			l1 = l1.Next
		} else {
			count++
			p.Next = l2
			p = p.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		p.Next = l1
		p = p.Next
	}
	if l2 != nil {
		p.Next = l2
		p = p.Next
	}
	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		p := head.Next
		head.Next = pre
		pre = head
		head = p
	}
	return pre
}

type ListNode struct {
	Val  int
	Next *ListNode
}
