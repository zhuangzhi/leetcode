package leetcode

import "sort"

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
	pairs := make([][2]int, 0, len(nums))
	for idx, val := range nums {
		pairs = append(pairs, [2]int{val, idx})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	r := len(pairs) - 1

	for l := 0; l < len(pairs); l++ {
		for l < r && pairs[l][0]+pairs[r][0] > target {
			r--
		}
		if l < r && pairs[l][0]+pairs[r][0] == target {
			return []int{pairs[l][1], pairs[r][1]}
		}
	}
	return []int{}
}
