package leetcode

// https://leetcode-cn.com/problems/merge-sorted-array/
// 88. 合并两个有序数组
//
func le88_merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for ; j >= 0; k-- {
		if i < 0 || nums1[i] <= nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
	}
}
