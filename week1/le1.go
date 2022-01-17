package leetcode

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/two-sum/
// 1. 两数之和
// 使用map key：value， value：idx,
// 首先构建map，然后计算每个值和9的差如果在map中存在就取得idx
func twoSum(nums []int, target int) []int {
	hashNum := map[int]int{}
	for idx, val := range nums {
		diff := target - val
		// 查看是否存在这个差值
		if p, ok := hashNum[diff]; ok {
			return []int{p, idx}
		}
		// 存入当前的val -> idx map pair
		hashNum[val] = idx
	}
	return []int{}
}

// 双指针方法, 先排序，再进行双指针扫描
func twoSum2(nums []int, target int) []int {
	in := Ints(nums)

	items := in.ToItems()
	items.Sort()

	r := items.Len() - 1

	for l, item := range items {
		for l < r && item.First+items[r].First > target {
			r--
		}
		for l < r && item.First+items[r].First == target {
			return []int{item.Second, items[r].Second}
		}
	}

	return []int{}
}
