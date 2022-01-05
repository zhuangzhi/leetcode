package leetcode

// https://leetcode-cn.com/problems/trapping-rain-water/
// 42. 接雨水
//

type rect42 struct {
	height, width int
}

func trap(height []int) int {
	stack := newanystack(len(height))
	ans := 0
	for _, height := range height {
		// 如果遇到更高的，开始一层一层，计算可以存的水，
		accumulatedWidth := 0

		for !stack.empty() && stack.top().(rect42).height < height {
			curr := stack.top().(rect42)
			accumulatedWidth += curr.width
			stack.pop()
			if !stack.empty() {
				minheight := minInt(stack.top().(rect42).height, height)
				ans += accumulatedWidth * (minheight - curr.height)
			}
		}
		stack.push(rect42{height, accumulatedWidth + 1})
	}
	return ans
}
