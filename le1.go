package leetcode

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
