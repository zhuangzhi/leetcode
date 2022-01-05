package leetcode

// https://leetcode-cn.com/problems/move-zeroes/
// 283. 移动零

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	for n, i := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[n] = nums[i]
			if n != i {
				nums[i] = 0
			}
			n++
		}
	}
}
