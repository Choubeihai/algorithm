package main

func main() {

}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var dummy = &ListNode{}
	dummy.Next = head
	pre := dummy
	tail := head
	for tail != nil {
		for i := 0; i < k-1 && tail != nil; i++ {
			tail = tail.Next
		}
		if tail == nil {
			break
		}
		next := tail.Next
		tail.Next = nil
		tail = reverseLinkList(head)
		tail, head = head, tail
		pre.Next = head
		tail.Next = next
		pre = tail
		head = next
		tail = next
	}
	return dummy.Next
}

func reverseLinkList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	for head != nil {
		cur := head
		next := cur.Next
		cur.Next = pre
		pre = cur
		head = next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
