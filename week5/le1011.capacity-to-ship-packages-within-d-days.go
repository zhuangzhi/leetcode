package week5

import "sort"

// https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days/
// 1011. 在 D 天内送达包裹的能力

func maxAndAccumulate(nums []int) (max, sum int) {
	for _, v := range nums {
		sum += v
		if max < v {
			max = v
		}
	}
	return
}

func shipWithinDaysII(weights []int, days int) int {
	left, right := maxAndAccumulate(weights)
	return left + sort.Search(right-left, func(i int) bool {
		return validate1011(weights, days, left+i)
	})
}

func shipWithinDays(weights []int, days int) int {
	left, right := maxAndAccumulate(weights)
	for left < right {
		mid := (left + right) / 2
		if validate1011(weights, days, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func validate1011(weights []int, days, cap int) bool {
	count := 1
	sum := 0
	for _, v := range weights {
		sum += v
		if sum > cap {
			count++
			sum = v
		}
	}
	return count <= days
}
