package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 利用分治法
func mergeKLists(lists []*ListNode) *ListNode {
	var n = len(lists)
	if n == 0 {
		return nil
	}
	var m = (n + 1) / 2
	var mid = n / 2
	for n != 1 {
		for i := 0; i < mid; i++ {
			lists[i] = merge(lists[i], lists[i+m])
		}
		n = m
		m = (n + 1) / 2
		mid = n / 2
	}
	return lists[0]
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
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
