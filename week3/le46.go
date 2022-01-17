package week3

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/permutations/
// 46. 全排列

func permute(nums []int) [][]int {
	n := len(nums)
	return backtrack(n, 0, make([][]int, 0, n*n), nums)
}

func backtrack(n, first int, ans [][]int, res Ints) [][]int {
	if n == first {
		return append(ans, res.Copy())
	}
	for i := first; i < n; i++ {
		res.Swap(first, i)
		ans = backtrack(n, first+1, ans, res)
		res.Swap(first, i)
	}
	return ans
}
