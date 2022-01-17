package leetcode

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/maximum-subarray/
// 53. 最大子数组和

// 前缀和s[i] = s[i-1] + nums[i-1]
// 任意字串和就是 s(l, r) = s[r] - s[l-1]
// nums 是可能有负数的, s[i]并不是单调的。
// 当前位置的最大子串和=s[i] - 之前的最小前缀和
// 这样我们还需要构造最小前缀和数组preMin[i] = min(s[i], preMin[i-1])
// ans = max(ans, s[i]-preMin[i-1])

func maxSubArray(nums []int) int {
	//利用前缀和，preMin
	n := len(nums)
	s := make([]int, n+1)
	preMin := make([]int, n+1)
	s[0] = 0
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + nums[i-1]
	}
	preMin[0] = s[0]
	for i := 1; i <= n; i++ {
		// 当前位置之前的最小的字串和， 有可能有nums可能有复数，s[i]不是单调增的。
		preMin[i] = MinInt(preMin[i-1], s[i])
	}

	ans := -10000000
	for i := 1; i <= n; i++ {
		ans = MaxInt(ans, s[i]-preMin[i-1])
	}
	return ans
}
