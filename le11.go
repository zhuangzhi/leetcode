package leetcode

// https://leetcode-cn.com/problems/container-with-most-water/
// 11. 盛最多水的容器

func maxArea(height []int) int {
	n := len(height)
	area := 0
	i, j := 0, n-1
	for i < j {
		area = maxInt(area, minInt(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return area
}
