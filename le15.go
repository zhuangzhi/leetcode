package leetcode

import "sort"

// https://leetcode-cn.com/problems/3sum/
// 15. 三数之和

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		jks := twoSumIII(nums[i+1:], -nums[i])
		for _, jk := range jks {
			ans = append(ans, []int{nums[i], jk[0], jk[1]})
		}
	}
	return ans
}

func twoSumIII(numbers []int, target int) [][]int {

	ans := [][]int{}
	r := len(numbers) - 1

	for l := 0; l < len(numbers); l++ {
		if l > 0 && numbers[l] == numbers[l-1] {
			continue
		}

		for l < r && numbers[l]+numbers[r] > target {
			r--
		}
		if l < r && numbers[l]+numbers[r] == target {
			ans = append(ans, []int{numbers[l], numbers[r]}) // 从1开始计数
		}
	}
	return ans
}
