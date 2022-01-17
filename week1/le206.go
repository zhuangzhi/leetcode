package leetcode

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/reverse-linked-list/
// 206. 反转链表

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	it, next, p2 := head, head.Next, head
	for next != nil {
		p2 = next.Next
		next.Next = it
		it = next
		next = p2
	}
	head.Next = nil
	return it
}
