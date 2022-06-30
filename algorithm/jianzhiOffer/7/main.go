package main

var m map[int]int

func buildTree(preorder []int, inorder []int) *TreeNode {
	m = make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		m[inorder[i]] = i
	}
	return dfs(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func dfs(preorder []int, preorderLeft, preorderRight int,
	inorder []int, inorderLeft, inorderRight int) *TreeNode {
	if preorderLeft > preorderRight {
		return nil
	}
	root := &TreeNode{Val: preorder[preorderLeft]}
	inIndex := m[preorder[preorderLeft]]
	leftLength := inIndex - inorderLeft
	rightLength := inorderRight - inIndex
	root.Left = dfs(preorder, preorderLeft+1, preorderLeft+leftLength, inorder, inorderLeft, inIndex-1)
	root.Right = dfs(preorder, preorderRight-rightLength+1, preorderRight, inorder, inIndex+1, inorderRight)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
