package main

/**
链表反转
*/

func main() {

}

func reverseLinkList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	for head != nil {
		cur := head
		next := cur.next
		cur.next = pre
		pre = cur
		head = next
	}
	return pre
}

type ListNode struct {
	val  int
	next *ListNode
}
