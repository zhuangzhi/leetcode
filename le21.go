package leetcode

// https://leetcode-cn.com/problems/merge-two-sorted-lists/
// 21. 合并两个有序链表

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
