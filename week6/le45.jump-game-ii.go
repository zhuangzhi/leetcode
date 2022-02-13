package week6

// https://leetcode-cn.com/problems/jump-game-ii/
// 45. 跳跃游戏 II
//
func jump(nums []int) int {
	now := 0
	target := len(nums) - 1
	ans := 0
	for now < target {
		//
		right := now + nums[now]
		if right >= target {
			return ans + 1
		}
		next, nextRight := now, right
		for i := now + 1; i <= right; i++ {
			if i+nums[i] > nextRight {
				nextRight = i + nums[i]
				next = i
			}
		}
		now = next
		ans++
	}
	return ans
}
