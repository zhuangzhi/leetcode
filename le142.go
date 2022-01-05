package leetcode

// https://leetcode-cn.com/problems/linked-list-cycle-ii/
// 142. 环形链表 II
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		if fast == slow {
			ptr := head
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
