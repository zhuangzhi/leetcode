package week5

// https://leetcode-cn.com/problems/reverse-pairs/
// 493. 翻转对
// 给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。
// 乱序问题，在一个乱序数组中统计满足特定大小关系的pair的数量，考虑基于归并排序求解。
// 分析：对这种既有位置的关系又有数值的关系，使用分治算法
func reversePairs(nums []int) int {
	return mergeSort493(nums, 0, len(nums)-1)
}

func mergeSort493(arr []int, l, r int) int {
	if l >= r {
		return 0
	}
	mid := (l + r) >> 1
	ans := mergeSort493(arr, l, mid)
	ans += mergeSort493(arr, mid+1, r)
	ans += calculate(arr, l, mid, r)
	merge493(arr, l, mid, r)
	return ans
}

func calculate(arr []int, l, mid, r int) int {
	// left [left, mid], right :[mid+1, right]
	ans := 0
	j := mid

	for i := l; i <= mid; i++ {
		for j < r && int64(arr[i]) > int64(arr[j+1])*2 {
			j++
		}
		ans += j - mid
	}
	return ans
}

func merge493(arr []int, l, mid, r int) {
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
