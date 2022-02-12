package week5

// https://leetcode-cn.com/problems/split-array-largest-sum/
// 410. 分割数组的最大值
func splitArray(nums []int, m int) int {
	max := -1
	sums := make([]int, len(nums)+1)
	sums[0] = 0
	for i, v := range nums {
		if max < v {
			max = v
		}
		sums[i+1] = sums[i] + v
	}
	sum := sums[len(nums)]
	left, right := max, sum
	for left < right {
		mid := (left + right) / 2
		if validate(sums, m, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func validate(sums []int, m, size int) bool {
	count := 0
	from := 0
	to := len(sums) - 1
	for count <= m {
		n := takeN(sums[from:], size)
		if n == 0 { // 没能拿出来，即size < nums[i]
			return false
		}
		count++
		from += n
		if from >= to { // 找到了最后
			break
		}
	}
	return count <= m
}

func takeN(sums []int, size int) int {
	left, right := 0, len(sums)-1
	for left < right {
		mid := (left + right + 1) / 2
		if sums[mid]-sums[0] <= size {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right
}
