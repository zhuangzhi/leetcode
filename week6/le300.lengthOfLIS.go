package week6

// https://leetcode-cn.com/problems/longest-increasing-subsequence/
// 300. 最长递增子序列
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
	}
	max, _ := maxInts(dp)
	return max
}

func maxInts(s []int) (max, id int) {
	max = s[0]
	for i, v := range s {
		if max < v {
			max = v
			id = i
		}
	}
	return
}
