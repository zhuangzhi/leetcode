package leetcode

import "fmt"

// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
// 25. K 个一组翻转链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

func reverseKGroup(head *ListNode, k int) *ListNode {
	to := takeGroupEnd(head, k)
	if to == nil {
		return head
	}
	nhead := reverse(head, to)
	head.Next = reverseKGroup(head.Next, k)
	return nhead
}

func reverse(head, to *ListNode) *ListNode {
	node, next := head, head.Next
	for node != to {
		t := next.Next
		next.Next = node
		node = next
		next = t
	}
	head.Next = next
	return node
}

func takeGroupEnd(head *ListNode, k int) *ListNode {
	k--
	for ; head != nil && k > 0; k-- {
		head = head.Next
	}
	return head
}
