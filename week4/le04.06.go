package week4

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/successor-lcci/
// 面试题 04.06. 后继者
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	return getNext(root, p.Val)
}

func getNext(root *TreeNode, val int) *TreeNode {
	var ans, curr *TreeNode
	curr = root
	for curr != nil {
		if curr.Val == val {
			if curr.Right != nil {
				ans = curr.Right
				for ans.Left != nil {
					ans = ans.Left
				}
			}
			break
		}
		if curr.Val < val {
			curr = curr.Right
		} else {
			// 记录历史
			if ans == nil || ans.Val > curr.Val {
				ans = curr
			}
			curr = curr.Left
		}
	}
	return ans
}
