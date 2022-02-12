package week5

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
// 215. 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	// 快排的本质
	return findKthLargestI(nums, len(nums)-k, 0, len(nums)-1)
}

func findKthLargestI(nums []int, k, l, r int) int {
	//
	if l >= r {
		return nums[l]
	}
	pivot := partitionII(nums, l, r)
	if pivot < k {
		return findKthLargestI(nums, k, pivot+1, r)
	} else {
		return findKthLargestI(nums, k, l, pivot)
	}
}

func partitionII(arr []int, l, r int) int {
	pivot := randInt(l, r)
	pivotVal := arr[pivot]
	for l <= r {
		for ; arr[l] < pivotVal; l++ {
		}
		for ; arr[r] > pivotVal; r-- {
		}
		if l == r {
			break
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	return r
}
