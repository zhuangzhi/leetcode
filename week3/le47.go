package week3

import (
	"sort"

	. "github.com/zhuangzhi/leetcode/util"
)

// permutation

func permuteUnique(nums []int) (ans [][]int) {
	n := len(nums)
	sort.Ints(nums)
	perm := Ints([]int{})
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, perm.Copy())
			return
		}
		for i, v := range nums {
			if vis[i] || (i > 0 && !vis[i-1] && v == nums[i-1]) {
				continue
			}
			perm = perm.Append(nums[i])
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm.Pop()
		}
	}

	backtrack(0)

	return
}

func backtrackUnique(n, first int, ans [][]int, res Ints) [][]int {
	if n == first {
		return append(ans, res.Copy())
	}
	for i := first; i < n; i++ {
		if res[first] == res[i] {
			continue
		}
		res.Swap(first, i)
		ans = backtrack(n, first+1, ans, res)
		res.Swap(first, i)
	}
	return ans
}
