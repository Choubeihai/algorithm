package main

/**
二叉搜索树的后序遍历
根节点 > 左子树
根节点 < 右子树
*/

func verifyPostorder(postorder []int) bool {
	return dfs(postorder, 0, len(postorder)-1)
}

func dfs(postorder []int, left, right int) bool {
	if left >= right {
		return true
	}
	var p = left
	for p <= right && postorder[p] < postorder[right] {
		p++
	}
	var m = p
	for p <= right && postorder[p] > postorder[right] {
		p++
	}
	if p != right {
		return false
	}
	return dfs(postorder, left, m-1) && dfs(postorder, m, right-1)
}
