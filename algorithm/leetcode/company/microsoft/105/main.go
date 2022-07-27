package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	var m = make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		m[inorder[i]] = i
	}
	n := len(preorder)
	return build(preorder, 0, n-1, inorder, 0, n-1, m)
}

func build(preorder []int, preLeft, preRight int, inorder []int, inLeft, inRight int, m map[int]int) *TreeNode {
	if preLeft > preRight {
		return nil
	}
	var root = &TreeNode{}
	root.Val = preorder[preLeft]
	leftLength := m[root.Val] - inLeft
	rightLength := inRight - m[root.Val]
	root.Left = build(preorder, preLeft+1, preLeft+leftLength, inorder, inLeft, m[root.Val]-1, m)
	root.Right = build(preorder, preRight-rightLength+1, preRight, inorder, m[root.Val]+1, inRight, m)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
