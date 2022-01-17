package util

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) ToString() string {
	if node == nil {
		return "nil"
	}
	return fmt.Sprint(node.Val)
}

func (node *ListNode) Goto(n int) *ListNode {
	if n == 0 {
		return node
	}
	if node.Next == nil {
		return nil
	}
	return node.Next.Goto(n)
}

func BuildListNodes(nums []int) *ListNode {
	var head ListNode
	node := &head
	for _, v := range nums {
		node.Next = &ListNode{Val: v}
		node = node.Next
	}
	return head.Next
}

func StoreListNodes(head *ListNode) []int {
	out := make([]int, 0, 512)
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}
