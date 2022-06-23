package _3

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p = head
	var value = p.Val

	for p != nil && p.Next != nil {
		if value == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
			value = p.Val
		}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
