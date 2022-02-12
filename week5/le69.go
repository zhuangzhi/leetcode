package week5

// https://leetcode-cn.com/problems/sqrtx/
// 69. x 的平方根
func mySqrt(x int) int {
	// 找最大的ans ans*ans <= x
	// 条件单调
	left, right := 0, x
	for left < right {

		mid := (left + right + 1) / 2 //正方向需要+1
		if mid <= x/mid {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right
}

func doubleSqrt(x float64) float64 {
	// 找最大的ans ans*ans <= x
	// 条件单调
	left, right := float64(0), x
	for (right - left) > 1e-7 {

		mid := (left + right) / 2 //正方向需要+1
		if mid*mid <= x {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}
