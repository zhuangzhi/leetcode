package leetcode

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
// 26. 删除有序数组中的重复项

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[n] = nums[i]
			n++
		}
	}
	return n
}
