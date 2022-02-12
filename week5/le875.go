package week5

import (
	"sort"
)

// https://leetcode-cn.com/problems/koko-eating-bananas/
// 875. 爱吃香蕉的珂珂
func minEatingSpeed(piles []int, h int) int {
	max, _ := maxAndAccumulate(piles)
	left, right := 1, max

	return left + sort.Search(right-left, func(x int) bool {
		return possible(piles, h, left+x)
	})
}

func possible(piles []int, h, k int) bool {
	time := 0
	for _, p := range piles {
		time += (p-1)/k + 1
	}
	return time <= h
}
