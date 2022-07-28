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
		root.Left.Next = root.Right
	}
	if root.Right != nil {
		if root.Next != nil {
			root.Right.Next = root.Next.Left
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
