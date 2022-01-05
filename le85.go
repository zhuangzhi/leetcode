package leetcode

// https://leetcode-cn.com/problems/maximal-rectangle/
// 85. 最大矩形
func maximalRectangle(matrix [][]byte) int {
	n := len(matrix)
	heights := make([][]int, n)
	for i := 0; i < n; i++ {
		heights[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				if i == 0 {
					heights[i][j] = 1
				} else {
					heights[i][j] = heights[i-1][j] + 1
				}
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = maxInt(ans, largestRectangleArea(heights[i]))
	}
	return ans
}
