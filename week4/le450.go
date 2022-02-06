package week4

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/delete-node-in-a-bst/
// 450. 删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		next := root.Right
		for next.Left != nil {
			next = next.Left
		}
		root.Val = next.Val
		root.Right = deleteNode(root.Right, next.Val)
		return root
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Left, key)
	}
	return root
}
