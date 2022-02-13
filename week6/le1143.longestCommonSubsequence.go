package week6

// https://leetcode-cn.com/problems/longest-common-subsequence/
// 1143. 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	f := make([][]int, m+1)
	f[0] = make([]int, n+1)
	for i, v1 := range text1 {
		f[i+1] = make([]int, n+1)

		for j, v2 := range text2 {
			if v1 == v2 {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = maxInt(f[i+1][j], f[i][j+1])
			}
		}
	}
	return f[m][n]
}
