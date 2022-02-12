package week5

import "sort"

// https://leetcode-cn.com/problems/merge-intervals/
// 56. 合并区间
func mergeInterval(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		x, y := intervals[i], intervals[j]
		return x[0] < y[0] || (x[0] == y[0] && x[1] < y[1])
	})
	ans := make([][]int, 0, len(intervals))
	start, farthest := -1, -1
	for _, v := range intervals {
		l, r := v[0], v[1]
		if l <= farthest {
			if farthest < r {
				farthest = r
			}
		} else {
			if farthest != -1 {
				ans = append(ans, []int{start, farthest})
			}
			start, farthest = l, r
		}
	}
	ans = append(ans, []int{start, farthest})
	return ans
}

func mergeInterval2(intervals [][]int) [][]int {
	type pair struct {
		f, s int
	}
	events := []pair{}
	for _, iv := range intervals {
		events = append(events, pair{iv[0], 1})
		events = append(events, pair{iv[1] + 1, -1})
	}
	sort.Slice(events, func(i, j int) bool {
		x, y := events[i], events[j]
		return x.f < y.f || (x.f == y.f && x.s < y.s)
	})

	ans := make([][]int, 0, len(intervals))
	covering := 0
	start := 0
	for _, e := range events {
		if covering == 0 {
			start = e.f
		}
		covering += e.s
		if covering == 0 {
			ans = append(ans, []int{start, e.f - 1})
		}
	}
	return ans
}
