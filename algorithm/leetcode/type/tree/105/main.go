package main

/**
根据前序序列和中序序列构建二叉树
思路：
前序序列先遍历根节点，然后遍历左子树，最后遍历右子树
中序序列先遍历左子树，在遍历根节点，最后遍历右子树
所以：关键是根据中序序列确定前序序列中左子树和右子树的长度
做法：
1. 先用map把中序序列的value和index存起来，value ——> index
2. 取前序序列的第一个value，构建成root节点
2. 这颗左子树的长度是:leftNum = m[value] - inStart
3. 这颗右子树的长度是:rightNum = inEnd- m[value]
4. 确定左子树的下标范围: 前序序列：[preStart+1, preStart+leftNum],中序序列：[inStart, m[value]-1]
4. 确定右子树的下标范围：前序序列：[preEnd-rightNum+1, preEnd], 中序序列：[m[value]+1, inEnd]

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	m := make(map[int]int, n)
	for i := 0; i < n; i++ {
		m[inorder[i]] = i
	}
	root := build(preorder, 0, n-1, inorder, 0, n-1, m)
	return root
}

func build(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int, m map[int]int) *TreeNode {
	// 出口
	if preStart > preEnd {
		return nil
	}
	// 出口
	if preStart == preEnd {
		return &TreeNode{Val: preorder[preStart]}
	}
	rootIndexInInorder := m[preorder[preStart]]
	leftNum := rootIndexInInorder - inStart
	rightNum := inEnd - rootIndexInInorder
	root := &TreeNode{Val: preorder[preStart]}
	root.Left = build(preorder, preStart+1, preStart+leftNum, inorder, inStart, rootIndexInInorder-1, m)
	root.Right = build(preorder, preEnd-rightNum+1, preEnd, inorder, rootIndexInInorder+1, inEnd, m)
	return root
}
