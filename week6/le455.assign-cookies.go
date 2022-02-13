package week6

import "sort"

// https://leetcode-cn.com/problems/assign-cookies/
// 455. 分发饼干
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	j := 0
	ans := 0
	for _, v := range g {
		for j < len(s) && s[j] < v {
			j++
		}
		if j == len(s) {
			break
		}
		j++
		ans++
	}
	return ans
}
