package leetcode

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/container-with-most-water/
// 11. 盛最多水的容器

func maxArea(height []int) int {
	n := len(height)
	area := 0
	i, j := 0, n-1
	for i < j {
		area = MaxInt(area, MinInt(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return area
}
