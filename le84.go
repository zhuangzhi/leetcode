package leetcode

//

import (
	"container/list"
)

// 单调栈的方法主要思路是通过计算矩形的面积替代一个一个柱子加起来算面积加速算法。也就是用乘法代替加法。
func largestRectangleArea(heights []int) int {
	// 填充一个0高度的柱子在最后，这样递增到最后的矩形都能被计算到
	heights = append(heights, 0)
	ans := 0
	s := list.New()
	// 循环所有的高度 _ 是idx。
	for _, height := range heights {
		accumulatedWidth := 0
		// 取出前一个矩形
		e := s.Front()

		// 如果前一个矩形存在并且高度大于当前的高度就要开始计算
		for e != nil {
			r, _ := e.Value.(Rect)
			if r.Height < height {
				break
			}
			accumulatedWidth += r.Width
			ans = maxInt(ans, r.Height*accumulatedWidth)
			s.Remove(e)
			e = s.Front()
		}
		// 以当前高度抹平之前比他高的矩形
		s.PushFront(Rect{accumulatedWidth + 1, height})
	}
	return ans
}

type Rect struct {
	Width, Height int
}
