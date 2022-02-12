package week5

import "math/rand"

// https://leetcode-cn.com/problems/sort-an-array/
// 912. 排序数组
func sortArray(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, l, mid, r int) {
	temp := make([]int, r-l+1)
	i, j := l, mid+1
	for k := range temp {
		if j > r || (i <= mid && arr[i] <= arr[j]) {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
	}
	for k := range temp {
		arr[l+k] = temp[k]
	}
}

func quickSort(arr []int, l, r int) {
	if l == r {
		return
	}
	pivot := partition(arr, l, r)
	quickSort(arr, l, pivot)
	quickSort(arr, pivot+1, r)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func partition(arr []int, l, r int) int {
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
