package main

func buildTree(inorder []int, postorder []int) *TreeNode {
	var m = make(map[int]int)
	var n = len(inorder)
	for i := 0; i < n; i++ {
		m[inorder[i]] = i
	}
	return build(inorder, 0, n-1, postorder, 0, n-1, m)
}

func build(inorder []int, inLeft, inRight int, postorder []int, postLeft, postRight int, m map[int]int) *TreeNode {
	if postLeft > postRight {
		return nil
	}
	var root = &TreeNode{}
	root.Val = postorder[postRight]
	leftLength := m[root.Val] - inLeft
	rightLength := inRight - m[root.Val]
	root.Left = build(inorder, inLeft, m[root.Val]-1, postorder, postLeft, postLeft+leftLength-1, m)
	root.Right = build(inorder, m[root.Val]+1, inRight, postorder, postRight-rightLength, postRight-1, m)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
