package leetcode

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/trapping-rain-water/
// 42. 接雨水
//

type rect42 struct {
	height, width int
}

func trap(height []int) int {
	stack := NewPairStack(len(height))
	ans := 0
	for _, height := range height {
		// 如果遇到更高的，开始一层一层，计算可以存的水，
		accumulatedWidth := 0

		for !stack.Empty() && stack.Top().First < height {
			curr := stack.Top()
			accumulatedWidth += curr.Second
			stack.Pop()
			if !stack.Empty() {
				minheight := MinInt(stack.Top().First, height)
				ans += accumulatedWidth * (minheight - curr.First)
			}
		}
		stack.Push(Pair{height, accumulatedWidth + 1})
	}
	return ans
}
