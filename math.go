package leetcode

func maxInt(l, r int) int {
	if l >= r {
		return l
	}
	return r
}

func minInt(l, r int) int {
	if l <= r {
		return l
	}
	return r
}

func EqualsMap(left, right map[string]int) bool {
	// if len(left) != len(right) {
	// 	return false
	// }
	for key, val := range left {
		if v1, ok := right[key]; ok {
			if val != v1 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
