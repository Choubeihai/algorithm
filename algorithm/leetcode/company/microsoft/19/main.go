package main

/**
思路：双指针，先让p1指针移动n次，然后同时移动p1，p2
细节：因为是要移除倒数第n个，返回移除后的整条链表，所以p2不能移动到第n个，需要移动第n-1个。
     所以，这里判定条件是p1.Next!=nil，但同时为了防止空指针，需要同时判断p1!=nil
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	p1 := dummy
	p2 := dummy
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	for p1 != nil && p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return dummy.Next

}

type ListNode struct {
	Val  int
	Next *ListNode
}
