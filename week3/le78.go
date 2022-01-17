package week3

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// 78. Subsets
func subsets(nums []int) [][]int {
	// 如果为空
	if len(nums) == 0 {
		return [][]int{{}}
	}
	curr := nums[0]
	// 计算后边子集的结果
	res := subsets(nums[1:])

	ans := IntsArray(res).Copy()
	ans.ComputeEach(func(i int, x Ints) Ints {
		return x.Copy().Append(curr)
	})
	// 合并子集的结果
	return ans.Append(res...)
}
