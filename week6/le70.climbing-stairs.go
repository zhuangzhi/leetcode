package week6

// https://leetcode-cn.com/problems/climbing-stairs/
// 70. 爬楼梯
func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
