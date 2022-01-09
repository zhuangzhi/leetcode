package leetcode

// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
// 167. 两数之和 II - 输入有序数组
func twoSumII(numbers []int, target int) []int {
	r := len(numbers) - 1

	for l := 0; l < len(numbers); l++ {
		for l < r && numbers[l]+numbers[r] > target {
			r--
		}
		if l < r && numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1} // 从1开始计数
		}
	}
	return []int{}
}
