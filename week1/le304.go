package leetcode

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{[][]int{}}
	}
	h, w := len(matrix), len(matrix[0])
	s := make([][]int, h+1)
	s[0] = make([]int, w+1)
	for i := 0; i < h; i++ {
		s[i+1] = make([]int, w+1)
		for j := 0; j < w; j++ {
			s[i+1][j+1] = s[i][j+1] + s[i+1][j] - s[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{
		sums: s,
	}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 > row2 {
		row1, row2 = row2, row1
	}
	if col1 > col2 {
		col1, col2 = col2, col1
	}
	sums := nm.sums
	return sums[row2+1][col2+1] - sums[row1][col2+1] - sums[row2+1][col1] + sums[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
