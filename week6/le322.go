package week6

// https://leetcode-cn.com/problems/coin-change/
// 322. 零钱兑换
// 模版题
func coinChange(coins []int, amount int) int {
	INF := int(1e9)
	opt := make([]int, amount+1)
	opt[0] = 0
	for i := 1; i <= amount; i++ {
		opt[i] = INF
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 { // 这枚硬币满足要求
				opt[i] = minInt(opt[i], opt[i-coins[j]]+1)
			}
		}
	}
	if opt[amount] >= INF {
		opt[amount] = -1
	}
	return opt[amount]
}

func minInt(i, j int) int {
	if i < j {
		return i
	}
	return j
}
