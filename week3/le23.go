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

type BinaryHeap struct {
	heap []*ListNode
}

func (heap *BinaryHeap) Empty() bool {
	return len(heap.heap) <= 1
}

func (heap *BinaryHeap) Root() *ListNode {
	return heap.heap[1]
}

func (heap *BinaryHeap) DeleteRoot() {
	end := len(heap.heap) - 1
	heap.heap[1] = heap.heap[end]
	heap.heap = heap.heap[:end]
	heap.heapifyDown(1)
}

func (heap *BinaryHeap) Insert(node *ListNode) {
	heap.heap = append(heap.heap, node)
	heap.heapifyUp(len(heap.heap) - 1)
}

func (heap *BinaryHeap) swap(i, j int) {
	heap.heap[i], heap.heap[j] = heap.heap[j], heap.heap[i]
}

func (heap *BinaryHeap) heapifyUp(p int) {
	for p > 1 {
		if heap.heap[p].Val < heap.heap[p/2].Val {
			heap.swap(p, p/2)
			p /= 2
		} else {
			break
		}
	}
}
func (heap *BinaryHeap) size() int {
	l := len(heap.heap) - 1
	if l <= 0 {
		return 0
	}
	return l
}

func (heap *BinaryHeap) heapifyDown(p int) {
	child := p * 2
	for child < heap.size() {
		other := p*2 + 1
		if other < heap.size() && heap.heap[other].Val < heap.heap[child].Val {
			child = other
		}
		if heap.heap[child].Val < heap.heap[p].Val {
			heap.swap(p, child)
			p = child
			child = p * 2
		} else {
			break
		}
	}

}
