package util

func Search(k, n int, f func(int) bool) int {
	// 后继型 用n表示无解
	left, right := k, n
	for left < right {
		mid := (left + right) / 2
		if f(mid) {
			right = mid //后继型，一侧包含
		} else {
			left = mid + 1 // 一侧不包含
		}
	}
	return right
}

func SearchPost(k, n int, f func(int) bool) int {
	// 前驱型 用-1表示无解
	left, right := k-1, n-1
	for left < right {
		mid := (left + right + 1) / 2
		// 为什么+1向上取整
		// 当有[left,right]相邻的时候 (left+right)/2 =left，如果只有分支1
		// 即nums[mid] <= target，陷入死循环
		if f(mid) {
			left = mid // 前驱型，一侧包含
		} else {
			right = mid - 1 // 一侧不包含
		}
	}
	return k - 1
}
