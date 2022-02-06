package week4

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/convert-bst-to-greater-tree/
// 538. 把二叉搜索树转换为累加树
func convertBST(root *TreeNode) *TreeNode {
	node, _ := convertBST_(root, 0)
	return node
}

func convertBST_(root *TreeNode, sum int) (*TreeNode, int) {
	if root == nil {
		return nil, sum
	}
	node := &TreeNode{}
	node.Right, sum = convertBST_(root.Right, sum)
	node.Val = sum + root.Val
	node.Left, sum = convertBST_(root.Left, node.Val)
	return node, sum
}
