package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			p.Next = list1
			p = p.Next
			list1 = list1.Next
		} else {
			p.Next = list2
			p = p.Next
			list2 = list2.Next
		}
	}
	for list1 != nil {
		p.Next = list1
		p = p.Next
		list1 = list1.Next
	}
	for list2 != nil {
		p.Next = list2
		p = p.Next
		list2 = list2.Next
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
