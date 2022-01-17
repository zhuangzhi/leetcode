package week3

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
// 94. 二叉树的中序遍历

func inorderTraversal(root *TreeNode) []int {
	v := Visitor{}
	v.visitInorder(root)
	return v.set
}

type Visitor struct {
	set []int
}

func (v *Visitor) visitInorder(node *TreeNode) {
	if node == nil {
		return
	}
	v.visitInorder(node.Left)
	v.set = append(v.set, node.Val)
	v.visitInorder(node.Right)
}

func (v *Visitor) visitPreorder(node *TreeNode) {
	if node == nil {
		return
	}
	v.set = append(v.set, node.Val)
	v.visitPreorder(node.Left)
	v.visitPreorder(node.Right)
}

func (v *Visitor) visitPostorder(node *TreeNode) {
	if node == nil {
		return
	}
	v.set = append(v.set, node.Val)
	v.visitPostorder(node.Left)
	v.visitPostorder(node.Right)
}
