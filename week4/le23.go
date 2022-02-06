package week4

import (
	. "github.com/zhuangzhi/leetcode/util"
)

func mergeKLists(lists []*ListNode) *ListNode {
	heap := HeapSlice(lists, func(i, j interface{}) bool {
		if i == nil || i.(*ListNode) == nil {
			return false
		}
		if j == nil || j.(*ListNode) == nil {
			return true
		}
		return i.(*ListNode).Val < j.(*ListNode).Val
	})
	head := &ListNode{}
	curr := head
	for heap.Len() > 0 {
		node := heap.Pop().(*ListNode)
		if node != nil {
			curr.Next = node
			curr = node
			if node.Next != nil {
				heap.Push(node.Next)
			}
		}
	}
	return head.Next
}
