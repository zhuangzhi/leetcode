package util

func MaxInt(s ...int) int {
	if len(s) == 0 {
		return int(-1e9)
	}
	max := s[0]
	for _, v := range s {
		if max < v {
			max = v
		}
	}
	return max
}

func MinInt(s ...int) int {
	if len(s) == 0 {
		return int(1e9)
	}
	min := s[0]
	for _, v := range s {
		if min > v {
			min = v
		}
	}
	return min
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
