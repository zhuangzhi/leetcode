package week6

// https://leetcode-cn.com/problems/triangle/
// 120. 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	i := len(triangle) - 2
	for ; i >= 0; i-- {
		s := triangle[i]
		b := triangle[i+1]
		for j := 0; j < len(s); j++ {
			if b[j] > b[j+1] {
				s[j] += b[j]
			} else {
				s[j] += b[j+1]
			}
		}
	}
	return triangle[0][0]
}
