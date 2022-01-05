package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLe21(t *testing.T) {
	units := []struct {
		l1, l2 []int
		Expect []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{}, []int{}},
	}
	for _, u := range units {
		l1 := newListNodes(u.l1)
		l2 := newListNodes(u.l2)
		head := mergeTwoLists(l1, l2)

		assert.Equal(t, u.Expect, listNodesNums(head))
	}
}

func TestLe26(t *testing.T) {
	units := []struct {
		Input  []int
		Expect []int
	}{
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{}, []int{}},
	}
	for _, u := range units {
		n := removeDuplicates(u.Input)
		assert.Equal(t, u.Expect, u.Input[:n])
	}
}

func newListNodes(nums []int) *ListNode {
	var head ListNode
	node := &head
	for _, v := range nums {
		node.Next = &ListNode{Val: v}
		node = node.Next
	}
	return head.Next
}

func listNodesNums(head *ListNode) []int {
	out := make([]int, 0, 512)
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

func TestLe25(t *testing.T) {
	// input := []int{1, 2, 3, 4, 5}
	units := []struct {
		nums   []int
		k      int
		expect []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
	}
	for _, u := range units {
		head := newListNodes(u.nums)
		out := reverseKGroup(head, u.k)
		nums := listNodesNums(out)
		assert.Equal(t, u.expect, nums, reverseKGroup(newListNodes(u.nums), u.k))
	}
}
func TestLe66(t *testing.T) {
	units := []struct {
		heights  []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{1, 9, 9}, []int{2, 0, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
	}

	for _, u := range units {
		assert.Equal(t, u.expected, plusOne(u.heights))
	}
}

func TestLe84(t *testing.T) {
	units := []struct {
		heights  []int
		expected int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 4}, 4},
	}

	for _, u := range units {
		assert.Equal(t, u.expected, largestRectangleArea(u.heights))
	}
}
func TestLe85(t *testing.T) {
	units := []struct {
		matrix   [][]byte
		expected int
	}{
		{
			[][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}},
			6,
		},
		{
			[][]byte{{'0'}},
			0,
		},
		{
			[][]byte{{'1'}},
			1,
		},
		{
			[][]byte{{'0', '0'}},
			0,
		},
	}

	for _, u := range units {
		assert.Equal(t, u.expected, maximalRectangle(u.matrix))
	}
}

func TestLe88(t *testing.T) {
	units := []struct {
		l1, l2 []int
		Expect []int
	}{
		{[]int{1, 2, 4, 0, 0, 0}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{}, []int{}},
	}
	for _, u := range units {
		le88_merge(u.l1, len(u.l1)-len(u.l2), u.l2, len(u.l2))
		assert.Equal(t, u.Expect, u.l1)
	}
}

func buildCycle(nums []int, pos int) *ListNode {
	head := newListNodes(nums)
	tail := head
	var node *ListNode
	for tail.Next != nil {
		if pos == 0 {
			node = tail
		}
		pos--
		tail = tail.Next
	}
	tail.Next = node
	return head
}

func TestLe141(t *testing.T) {
	// input := []int{1, 2, 3, 4, 5}
	units := []struct {
		nums     []int
		pos      int
		expected bool
	}{
		{[]int{1, 2, 3, -4}, 1, true},
		{[]int{1, 2, 3, -4, 5, 6, 7}, 3, true},
	}

	for _, u := range units {
		head := buildCycle(u.nums, u.pos)
		assert.Equal(t, u.expected, hasCycle(head))
	}
}

func TestLe142(t *testing.T) {
	// input := []int{1, 2, 3, 4, 5}
	units := []struct {
		nums     []int
		pos      int
		expected string
	}{
		{[]int{1, 2, 3, -4}, 1, "2"},
		{[]int{1, 2, 3, -4, 5, 6, 7}, 3, "-4"},
		{[]int{1, 2, 3, -4, 5, 6, 7}, -1, "nil"},
	}

	for _, u := range units {
		head := buildCycle(u.nums, u.pos)
		assert.Equal(t, u.expected, detectCycle(head).ToString())
	}
}

func TestLe206(t *testing.T) {
	// input := []int{1, 2, 3, 4, 5}
	units := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, u := range units {
		assert.Equal(t, u.expected, listNodesNums(reverseList(newListNodes(u.nums))))
	}
}

func TestLe283(t *testing.T) {
	units := []struct {
		nums     []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}

	for _, u := range units {
		moveZeroes(u.nums)
		assert.Equal(t, u.expected, u.nums)
	}
}
func TestLe641(t *testing.T) {
	a := assert.New(t)
	que := Constructor(2)
	a.Equal(0, que.size)
	que.InsertFront(1)
	que.InsertLast(2)
	a.Equal(2, que.size)
	a.Equal(1, que.GetFront())
	que.DeleteFront()
	a.Equal(1, que.size)
	a.Equal(2, que.GetFront())
}
