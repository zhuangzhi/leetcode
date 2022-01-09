package leetcode

// https://leetcode-cn.com/problems/count-number-of-nice-subarrays/
// 1248. 统计「优美子数组」
func numberOfSubarrays(nums []int, k int) int {
	// 所以数%2, 奇数为1, 偶数为 0
	// 求 sum(l, r) == k
	// S[r] - S[l-1] == k  (l<=r)
	n := len(nums) + 1
	s := make([]int, n)
	s[0] = 0
	for idx, v := range nums {
		s[idx+1] = s[idx] + v%2
	}
	//     9,8,6,7
	// ->  1,0,0,1  nums[i] % 2
	// ->0,1,1,1,2  s[i+1] = s[i] + nums[i] %2
	// ->count[1]++, ++, ++ = 3 (三组字串奇数和为1)
	// ->1,3,1,
	ans := 0
	counts := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i]-k >= 0 {
			ans += counts[s[i]-k]
		}
		counts[s[i]]++
	}
	return ans
}
