package week3

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// 236. 二叉树的最近公共祖先

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ans, _, _ := dfsCA(root, p, q)
	return ans
}

func dfsCA(root, p, q *TreeNode) (ans *TreeNode, hasP bool, hasQ bool) {
	if root == nil {
		return nil, false, false
	}
	lp, l1, l2 := dfsCA(root.Left, p, q)
	rp, r1, r2 := dfsCA(root.Right, p, q)
	if lp != nil {
		return lp, true, true
	}
	if rp != nil {
		return lp, true, true
	}
	hasP = l1 || r1 || root == p
	hasQ = l2 || r2 || root == q
	if hasP && hasQ {
		ans = root
	}
	return
}
