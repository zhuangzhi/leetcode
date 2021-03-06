package week4

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
// 701. 二叉搜索树中的插入操作
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
