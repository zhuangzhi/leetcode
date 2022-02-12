package week5

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	// 开始位置，第一个>=target的数
	// 结束为止，最后一个<=target的数

	// 后继型 用n表示无解
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid //后继型，一侧包含
		} else {
			left = mid + 1 // 一侧不包含
		}
	}
	s := right

	// 前驱型 用-1表示无解
	left, right = -1, len(nums)-1
	for left < right {
		mid := (left + right + 1) / 2
		// 为什么+1向上取整
		// 当有[left,right]相邻的时候 (left+right)/2 =left，如果只有分支1
		// 即nums[mid] <= target，陷入死循环
		if nums[mid] <= target {
			left = mid // 前驱型，一侧包含
		} else {
			right = mid - 1 // 一侧不包含
		}
	}
	t := right
	if s > t {
		return []int{-1, -1}
	}
	return []int{s, t}
}
