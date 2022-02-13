package week6

// https://leetcode-cn.com/problems/maximum-product-subarray/
// 152. 乘积最大子数组
func maxProduct(nums []int) int {
	n := len(nums)
	fmax := make([]int, n)
	fmin := make([]int, n)
	fmax[0] = nums[0]
	fmin[0] = nums[0]
	for i := 1; i < n; i++ {
		a, b, c := fmax[i-1]*nums[i], fmin[i-1]*nums[i], nums[i]
		fmax[i] = maxInt(a, b, c)
		fmin[i] = minInt(a, b, c)

	}
	max := maxInt(fmax...)
	return max
}
