package week3

import (
	. "github.com/zhuangzhi/leetcode/util"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder)
}
func build(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	mid := -1
	// mid: root在inorder中的位置
	// l1 -> mid-1 左子树中序
	// mid+1>l2 右子树的中序
	for idx, v := range inorder {
		if v == root.Val {
			mid = idx
			break
		}
	}
	if mid >= 0 {
		root.Left = build(preorder[1:mid+1], inorder[:mid])
		root.Right = build(preorder[mid+1:], inorder[mid+1:])
	}
	return root
}
