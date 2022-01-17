package week3

import (
	. "github.com/zhuangzhi/leetcode/util"
)

func combine(n int, k int) [][]int {
	// 不符合条件的返回空集合
	if n <= 0 || k <= 0 || n < k {
		return [][]int{}
	}
	// k = 1， 返回 1到n的n个集合
	if k == 1 {
		r := make([][]int, n)
		for i := 0; i < n; i++ {
			r[i] = []int{i + 1}
		}
		return r
	}
	// n = k， 返回 1到n组成的一个集合
	if n == k {
		return [][]int{IntsFromRange(1, n+1)}
	}

	// 计算不含当前值的k-1集合，之后再加上当前值
	pre := combine(n-1, k-1)
	// 计算不含当前值的k-1集合
	exceptMe := combine(n-1, k)

	// ans := make([][]int, 0, len(exceptMe)+len(pre))

	ans := IntsArray(pre)
	ans.ComputeEach(func(i int, x Ints) Ints {
		return x.Append(n)
	})
	// 合并包含，不包含当前值的两个集合
	return ans.Append(exceptMe...)
}
