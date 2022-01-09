package leetcode

// https://leetcode-cn.com/problems/number-of-submatrices-that-sum-to-target/
// 1074. 元素和为目标值的子矩阵数量

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	n := len(matrix)
	w := len(matrix[0])
	ans := 0
	for i := 0; i < n; i++ {
		sum := make([]int, w)
		for j := i; j < n; j++ {
			for c := 0; c < w; c++ {
				sum[c] += matrix[j][c]
			}
			ans += subarraySum(sum, target)
		}
	}
	return ans
}
