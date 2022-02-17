package week6

import "sort"

// https://leetcode-cn.com/problems/minimum-initial-energy-to-finish-tasks/
// 1665. 完成所有任务的最少初始能量
// 必须有一个完整的证明
// actual[p] - mini[p] <  actual[q] - mini[q]
// 任意位置p,q需要的能量 max(max(minimum[q], S+actual[q])+actual[p], minimum[p])
//  先q后p：max(max(minimum[p], S+actual[p])+actual[q], minimum[q])
// max(max(minimum[q], S+actual[p])+actual[p], minimum[p]) =>
// p,q: max(minimum[q] + actual[p], S+actual[p]+actual[q], minimum[p])
// q,p: max(minimum[p] + actual[q], S+actual[p]+actual[q], minimum[q])
// 比较上下两个大小，S+actual[p]+actual[q] 去掉
// p,q: max(minimum[q] + actual[p], minimum[q])
// q,p: max(minimum[p] + actual[q], minimum[q])
// 则若p,q较好：max(minimum[q] + actual[p], minimum[p])<max(minimum[p] + actual[q], minimum[q])
// 因为minimum[p] + actual[q] > minimum[p]， 所以这两项去掉
// 则若p,q较好：minimum[q] + actual[p] < minimum[p] + actual[q]
//  即：actual[p]- minimum[p] <  actual[q] - minimum[q]
// 因此可以按照 actual[i]- minimum[i]进行排序
func minimumEffort(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		p, q := tasks[i], tasks[j]
		return p[0]-p[1] > q[0]-q[1]
	})
	ans := 0
	for _, p := range tasks {
		ans = maxInt(p[0]+ans, p[1])
	}
	return ans
}
