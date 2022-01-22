package week4

// https://leetcode-cn.com/problems/n-queens/
// 51. N 皇后
func solveNQueens(n int) [][]string {
	ans := [][]string{}
	colused := make([]bool, n)
	digs1 := make([]bool, n*2) //  `/` 正对角线: row+col 固定值
	digs2 := make([]bool, n*2) //  `\` 反对角线: row-col 固定值，为了使用数组取 n+row-col >= 0
	var dfs func(row int)
	r := make([]int, n)
	dfs = func(row int) {
		if row == n {
			// build ans from r
			ans1 := []string{}
			for _, pos := range r {
				s := make([]rune, n)
				for i := range s {
					s[i] = '.'
					if i == pos {
						s[i] = 'Q'
					}
				}
				ans1 = append(ans1, string(s))
			}
			ans = append(ans, ans1)
			return
		}
		for col, used := range colused {
			if !used && !digs1[row+col] && !digs2[n+row-col] {
				//
				colused[col] = true
				digs1[row+col] = true
				digs2[n+row-col] = true
				r[row] = col
				dfs(row + 1)
				colused[col] = false
				digs1[row+col] = false
				digs2[n+row-col] = false
				r[row] = 0
			}
		}
	}
	dfs(0)
	return ans
}
