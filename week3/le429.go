package week3

// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal
// 429. N 叉树的层序遍历
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := make([][]int, 0, 256)
	if root == nil {
		return result
	}
	nodes := make([]*Node, 0, 256)
	children := make([]*Node, 0, 256)
	nodes = append(nodes, root)
	for len(nodes) != 0 {
		level := make([]int, 0, 32)
		//
		for _, node := range nodes {
			level = append(level, node.Val)
			children = append(children, node.Children...)
		}
		result = append(result, level)
		level = level[:0]
		nodes = append(nodes[:0], children...)
		children = children[:0]
	}
	return result
}
