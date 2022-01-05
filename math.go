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
