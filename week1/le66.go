package leetcode

// https://leetcode-cn.com/problems/plus-one/
// 66. 加一

func plusOne(digits []int) []int {
	last := len(digits) - 1
	// from last to fist find a NONE 9 number, otherwise + 1.
	for i := last; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}
	// All numbers is 9. return a new array start with 1.
	out := make([]int, len(digits)+1)
	out[0] = 1
	return out
}
