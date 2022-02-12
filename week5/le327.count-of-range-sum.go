package week5

import "fmt"

// https://leetcode-cn.com/problems/count-of-range-sum/
// 327. 区间和的个数
func countRangeSum(nums []int, lower int, upper int) int {
	sums := accumulate(nums)
	// fmt.Println(nums, sums)
	// return countRangeSumIn(sums, lower, upper, 0, len(sums)-1)
	var mergeCount func([]int) int
	mergeCount = func(arr []int) int {
		n := len(arr)
		if n <= 1 {
			return 0
		}
		n1 := append([]int(nil), arr[:n/2]...)
		n2 := append([]int(nil), arr[n/2:]...)
		cnt := mergeCount(n1) + mergeCount(n2)
		//
		l, r := 0, 0
		for _, v := range n1 {
			for l < len(n2) && n2[l]-v < lower {
				l++
			}
			for r < len(n2) && n2[r]-v <= upper {
				r++
			}
			cnt += r - l
		}
		p1, p2 := 0, 0
		for i := range arr {
			if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
				arr[i] = n1[p1]
				p1++
			} else {
				arr[i] = n2[p2]
				p2++
			}
		}
		return cnt
	}
	return mergeCount(sums)
}

func countRangeSumIn(sums []int, lower int, upper int, l, r int) int {
	if l >= r-1 {
		return 0
	}
	mid := (l + r + 1) / 2
	count := 0
	count += countRangeSumIn(sums, lower, upper, l, mid)
	count += countRangeSumIn(sums, lower, upper, mid, r)
	count += countMergeRangeSumIn(sums, lower, upper, l, mid, r)
	merge(sums, l, mid, r)
	fmt.Println(sums)
	return count
}

func countMergeRangeSumIn(sums []int, lower int, upper int, l, mid, r int) int {
	if l >= r-1 {
		return 0
	}
	count := 0
	// i, j := l, mid+1
	for i := l; i <= mid; i++ {
		for j := mid + 1; j <= r; j++ {
			fmt.Println("loop", i, mid, j)
			if sums[j]-sums[i] >= lower && sums[j]-sums[i] <= upper {
				fmt.Println(sums[j], sums[i], lower, upper)
				count++
			}
		}
	}
	return count
}

func accumulate(nums []int) []int {
	sums := make([]int, len(nums)+1)
	for i, v := range nums {
		sums[i+1] = sums[i] + v
	}
	return sums
}
