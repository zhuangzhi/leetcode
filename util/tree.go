package util

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) PreVisit(s *Ints) {
	if node == nil {
		return
	}
	s.Append(node.Val)
	node.Left.PreVisit(s)
	node.Right.PreVisit(s)
}

func DeSerializeTree(s string) *TreeNode {
	return nil
}

func (node *TreeNode) Serialize() string {
	ss := []string{}
	ss = node.store(ss)
	return strings.Join(ss, ",")
}

func (node *TreeNode) store(ss []string) []string {
	if node == nil {
		ss = append(ss, "null")
		return ss
	}
	ss = append(ss, strconv.Itoa(node.Val))
	ss = node.Left.store(ss)
	ss = node.Right.store(ss)
	return ss
}
