package week5

// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/
// 153. 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	// 和 nums[n-1]比较 nums[i] <= nums[n-1] 34512->00011
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] <= nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[right]
}
