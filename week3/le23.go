package week3

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/merge-k-sorted-lists/
// 23. 合并K个升序链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return mergeTwoLists(lists[0], lists[1])

	}
	split := len(lists) / 2
	return mergeTwoLists(mergeKLists(lists[0:split]), mergeKLists(lists[split:]))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	it := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			it.Next = list1
			it = list1
			list1 = list1.Next
		} else {
			it.Next = list2
			it = list2
			list2 = list2.Next
		}
	}
	if list1 == nil {
		it.Next = list2
	} else {
		it.Next = list1
	}
	return head.Next
}
