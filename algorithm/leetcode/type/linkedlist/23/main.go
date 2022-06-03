package main

func main() {

}

// 利用优先队列，基于堆排序
func mergeKLists2(lists []*ListNode) *ListNode {
	return nil
}

// 利用分治法
func mergeKLists(lists []*ListNode) *ListNode {
	var n = len(lists)
	for n > 1 {
		mid := (n + 1) / 2
		for i := 0; i < n/2; i++ {
			lists[i] = merge(lists[i], lists[i+mid])
		}
		n = mid
	}
	return lists[0]
}

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
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
