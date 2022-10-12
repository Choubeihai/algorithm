package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	for n != 1 {
		k := (n + 1) / 2
		for i := 0; i < n/2; i++ {
			lists[i] = merge(lists[i], lists[i+k])
		}
		n = k
	}
	return lists[0]
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
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
