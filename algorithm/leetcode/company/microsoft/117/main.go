package main

func connect(root *Node) *Node {
	dfs(root)
	return root
}

func dfs(root *Node) {
	if root == nil {
		return
	}
	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else if root.Next != nil {
			p := root.Next
			for p != nil && p.Left == nil && p.Right == nil {
				p = p.Next
			}
			if p != nil {
				if p.Left != nil {
					root.Left.Next = p.Left
				} else {
					root.Left.Next = p.Right
				}
			}
		}
	}
	if root.Right != nil {
		if root.Next != nil {
			p := root.Next
			for p != nil && p.Left == nil && p.Right == nil {
				p = p.Next
			}
			if p != nil {
				if p.Left != nil {
					root.Right.Next = p.Left
				} else {
					root.Right.Next = p.Right
				}
			}
		}
	}
	dfs(root.Left)
	dfs(root.Right)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
