package leetcode

// https://leetcode-cn.com/problems/subarray-sum-equals-k/
// 560. 和为 K 的子数组
func subarraySum(nums []int, k int) int {
	n := len(nums)
	s := make([]int, n+1)
	s[0] = 0
	smap := make(map[int]int, n+1)
	smap[0] = 1
	ans := 0
	pre := 0
	for _, v := range nums {
		pre += v
		if v, ok := smap[pre-k]; ok {
			ans += v
		}
		smap[pre]++
	}

	return ans
}
