package leetcode

import (
	"sort"

	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/3sum/
// 15. 三数之和

func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		jks := twoSumIII(nums[i+1:], -nums[i])
		for _, jk := range jks {
			ans = append(ans, []int{nums[i], jk[0], jk[1]})
		}
	}
	return ans
}

func twoSumIII(numbers []int, target int) [][]int {
	ans := [][]int{}
	r := len(numbers) - 1

	for l := 0; l < len(numbers); l++ {
		if l > 0 && numbers[l] == numbers[l-1] {
			continue
		}

		for l < r && numbers[l]+numbers[r] > target {
			r--
		}
		if l < r && numbers[l]+numbers[r] == target {
			ans = append(ans, []int{numbers[l], numbers[r]}) // 从1开始计数
		}
	}
	return ans
}

type Predicate struct {
}

func twoSumIIII(numbers []int, target int) [][]int {
	nums := Ints(numbers)
	ans := IntsArray([][]int{})
	r := len(numbers) - 1

	nums.ForEach(func(l, first int) {
		if nums.ItemEqualWithPre(l) {
			return // Skip duplicated values
		}
		next := 0
		r, next = nums.ReversLoopUntil(r, func(idxR, valR int) bool {
			return l < idxR && first+valR > target
		})
		if l < r && first+next == target {
			ans = ans.Append([]int{first, next}) // 从1开始计数
		}
	})

	return ans
}

func nSum(nums []int, n, target int) [][]int {
	if n == 2 {
		return twoSumIIII(nums, target)
	}
	in := Ints(nums)
	in.Sort()
	ans := IntsArray(nil)
	in.ForEach(func(idx, val int) {
		if idx > 0 && val == in[idx-1] {
			return
		}
		jks := nSum(in[idx+1:], n-1, target-val)
		ans = ans.Append(IntsArray(jks).ComputeEach(func(i int, s Ints) Ints {
			return s.Append(val)
		})...)
	})
	return ans
}
