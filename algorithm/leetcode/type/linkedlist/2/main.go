package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	bit := 0
	dummy := &ListNode{}
	pre := dummy
	for l1 != nil && l2 != nil {
		value := l1.Val + l2.Val + bit
		val := value % 10
		bit = value / 10
		node := &ListNode{Val: val}
		pre.Next = node
		pre = pre.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		value := l1.Val + bit
		val := value % 10
		bit = value / 10
		node := &ListNode{Val: val}
		pre.Next = node
		pre = pre.Next
		l1 = l1.Next
	}

	for l2 != nil {
		value := l2.Val + bit
		val := value % 10
		bit = value / 10
		node := &ListNode{Val: val}
		pre.Next = node
		pre = pre.Next
		l2 = l2.Next
	}
	if bit != 0 {
		node := &ListNode{Val: bit}
		pre.Next = node
	}
	return dummy.Next
}
