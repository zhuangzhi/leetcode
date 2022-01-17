package week3

import (
	. "github.com/zhuangzhi/leetcode/util"
)

func buildTree2(inorder []int, postorder []int) *TreeNode {
	n := len(postorder) - 1
	if n < 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[n]}
	mid := -1
	for i := n; i >= 0; i-- {
		if inorder[i] == root.Val {
			mid = i
			break
		}
	}
	if mid >= 0 {
		root.Left = buildTree2(inorder[0:mid], postorder[0:mid])
		root.Right = buildTree2(inorder[mid+1:], postorder[mid:n])
	}
	return root
}
