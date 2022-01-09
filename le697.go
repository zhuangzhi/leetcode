package leetcode

// https://leetcode-cn.com/problems/degree-of-an-array/
// 697. 数组的度

func findShortestSubArray(nums []int) int {
	type unit struct {
		left, right, degree int
	}
	degreeMap := make(map[int]*unit)
	degree := 0
	for idx, val := range nums {
		if u, ok := degreeMap[val]; ok {
			u.right = idx
			u.degree++
			degree = maxInt(degree, u.degree)
		} else {
			degreeMap[val] = &unit{left: idx, right: idx, degree: 1}
			degree = maxInt(degree, 1)
		}
	}
	ans := len(nums)
	for _, u := range degreeMap {
		if u.degree == degree {
			ans = minInt(ans, u.right-u.left+1)
		}
	}
	return ans
}
