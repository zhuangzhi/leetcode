package leetcode

// https://leetcode-cn.com/problems/corporate-flight-bookings/
// 1109. 航班预订统计

func corpFlightBookings(bookings [][]int, n int) []int {
	// 0 ~ n+2, 因为 n -> 1 ~ n
	// flight -> [1,n], delta[0, n+2)
	delta := make([]int, n+2)
	for _, v := range bookings {
		first, last, val := v[0], v[1], v[2]
		delta[first] += val
		delta[last+1] -= val
	}
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + delta[i]
	}
	// remove 0, 没有第0个航班
	return sum[1:]
}
