package main

/**
fast指针每次比slow多移动1步，如果有环，fast一定能追上slow
*/

func hasCycle(head *ListNode) bool {
	var slow = head
	var fast = head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}
