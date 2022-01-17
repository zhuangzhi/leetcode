package leetcode

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/sliding-window-maximum/
// 239. 滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	max := NewRing(k) // Store cadidates in ring, with max num at top.
	ans := make([]int, 0, len(nums)-k+1)
	for idx, num := range nums {
		// Delete left max idx out of window (k)
		// idx 4, k=2, remove 0,1,2, keep(3,4)
		for !max.IsEmpty() && max.GetFront().(int) <= idx-k {
			max.DeleteFront()
		}
		// Delete
		for !max.IsEmpty() && nums[max.GetRear().(int)] <= num {
			max.DeleteLast()
		}
		max.InsertLast(idx)
		// k = 2, idx >= 1 (0,1)
		if idx >= k-1 {
			ans = append(ans, nums[max.GetFront().(int)])
		}
	}
	return ans
}
