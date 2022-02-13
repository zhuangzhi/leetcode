package week6

// https://leetcode-cn.com/problems/lemonade-change/
// 860. 柠檬水找零
func lemonadeChange(bills []int) bool {
	coins := map[int]int{}
	exchange := func(amount int) bool {
		for _, coin := range []int{20, 10, 5} {
			for amount >= coin && coins[coin] > 0 {
				coins[coin]--
				amount -= coin
			}
		}
		return amount == 0
	}
	for _, v := range bills {
		coins[v]++
		if !exchange(v - 5) {
			return false
		}
	}
	return true
}
