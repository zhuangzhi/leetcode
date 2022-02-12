package week5

// https://leetcode-cn.com/problems/minimum-number-of-days-to-make-m-bouquets/
// 1482. 制作 m 束花所需的最少天数
func minDays(bloomDay []int, m int, k int) int {
	// 00001111 是否满足要求是单调的

	latestBloomDay := 0
	for _, v := range bloomDay {
		if latestBloomDay < v {
			latestBloomDay = v
		}
	}

	left, right := 0, latestBloomDay+1 // 代表无解

	for left < right {
		day := (left + right) / 2
		if validateOnDay(bloomDay, day, m, k) {
			right = day
		} else {
			left = day + 1
		}
	}
	if right == latestBloomDay+1 {
		return -1
	}
	return right
}

func validateOnDay(bloomDay []int, now, m, k int) bool {
	bouquet := 0     // 几个花束
	consecutive := 0 // 连续多少个
	for _, day := range bloomDay {
		if day <= now {
			consecutive++
			if consecutive == k {
				bouquet++
				consecutive = 0
			}
		} else {
			consecutive = 0
		}
	}
	return bouquet >= m
}
