package main

func main() {

}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var dummy = &ListNode{}
	dummy.Next = head
	var first, end = dummy, dummy
	for end != nil {
		for i := 0; i < k; i++ {
			if end == nil {
				break
			} else {
				end = end.Next
			}
		}
		if end == nil {
			break
		}
		p := end.Next
		end.Next = nil
		p2 := first.Next
		first.Next = reverse(first.Next)
		p2.Next = p
		first = p2
		end = p2
	}
	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		p := head.Next
		head.Next = pre
		pre = head
		head = p
	}
	return pre
}

type ListNode struct {
	Val  int
	Next *ListNode
}
