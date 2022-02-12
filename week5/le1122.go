package week5

import "sort"

// https://leetcode-cn.com/problems/relative-sort-array/
// 1122. 数组的相对排序
// arr2[i] uniq, < 1000
// 0< arr1[i], arr2[i] <= 1000,
// 1< len(arr1), len(arr2) <= 1000
func relativeSortArray(arr1 []int, arr2 []int) []int {
	upper := 0
	for _, v := range arr1 {
		if upper < v {
			upper = v
		}
	}
	mark := make([]int, upper+1)
	for id, v := range arr2 {
		mark[v] = id + 1
	}

	sort.Slice(arr1, func(i, j int) bool {
		x, y := arr1[i], arr1[j]
		mx, my := mark[x], mark[y]
		if mx == 0 {
			mx = x + 2000
		}
		if my == 0 {
			my = y + 2000
		}
		return mx < my
	})
	return arr1
}

func count(arr1 []int, arr2 []int) []int {
	upper := 0
	for _, v := range arr1 {
		if upper < v {
			upper = v
		}
	}
	count := make([]int, upper+1)
	for _, v := range arr1 {
		count[v]++
	}
	ans := make([]int, 0, len(arr1))
	for _, v := range arr2 {
		for ; count[v] > 0; count[v]-- {
			ans = append(ans, v)
		}
	}
	for v, freq := range count {
		for ; freq > 0; freq-- {
			ans = append(ans, v)
		}
	}
	return ans
}
