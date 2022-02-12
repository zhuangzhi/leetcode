package week5

// https://leetcode-cn.com/problems/find-peak-element/
// 162. 寻找峰值
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		lmid := (left + right) / 2
		rmid := lmid + 1
		if nums[lmid] <= nums[rmid] {
			left = lmid + 1
		} else {
			right = rmid - 1
		}
	}
	return right
}
