package week6

// https://leetcode-cn.com/problems/number-of-longest-increasing-subsequence/
// 673. 最长递增子序列的个数
func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	cnts := make([]int, len(nums))
	maxLen, ans := 0, 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		cnts[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// cnts[i]++
				if dp[i] == dp[j]+1 {
					cnts[i] += cnts[j]
				} else if dp[i] < dp[j]+1 {
					cnts[i] = cnts[j]
					dp[i] = dp[j] + 1
				}
				// dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
		if maxLen == dp[i] {
			ans += cnts[i]
		}
		if maxLen < dp[i] {
			maxLen = dp[i]
			ans = cnts[i]
		}
	}
	return ans
}
