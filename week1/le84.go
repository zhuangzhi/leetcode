package leetcode

import (
	"fmt"

	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
// 84. Largest Rectangle in Histogram
// 84. 柱状图中最大的矩形

// 单调栈的方法主要思路是通过计算矩形的面积替代一个一个柱子加起来算面积加速算法。也就是用乘法代替加法。
func largestRectangleArea(heights []int) int {
	// 填充一个0高度的柱子在最后，这样递增到最后的矩形都能被计算到
	heights = append(heights, 0)
	fmt.Println(heights)
	ans := 0
	s := NewPairStack(len(heights))
	// s := list.New()
	// 循环所有的高度 _ 是idx。
	for _, height := range heights {
		accumulatedWidth := 0
		// 如果前一个矩形存在并且高度大于当前的高度就要开始计算
		for e := s.Top(); !s.Empty() && e.First >= height; e = s.Top() {
			// e := s.Top()
			accumulatedWidth += e.Second
			ans = MaxInt(ans, e.First*accumulatedWidth)
			s.Pop()
		}
		// 以当前高度抹平之前比他高的矩形
		s.Append(height, accumulatedWidth+1)
	}
	return ans
}

type Rect struct {
	Width, Height int
}
