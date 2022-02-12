package week5

// https://leetcode-cn.com/problems/binary-search/
// 704. 二分查找

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		}
	}
	return -1
}
