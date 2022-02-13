package week6

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
// 122. 买卖股票的最佳时机 II
func maxProfit(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		v := prices[i] - prices[i-1]
		if v > 0 {
			ans += v
		}
	}
	return ans
}
