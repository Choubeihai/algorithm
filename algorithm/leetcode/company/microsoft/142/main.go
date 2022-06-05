package main

/**
fast指针每次比slow多移动1步，如果有环，fast一定能追上slow
如果有环，fast指针一定比slow都走一个环的距离
*/

func detectCycle(head *ListNode) *ListNode {
	var fast = head
	var slow = head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow = head
			for slow != fast {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
