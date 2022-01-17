package leetcode

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/
// 30. 串联所有单词的子串
// 滑动窗口 0->len(s)-k. (k = len(words) * len(words[0])
// split substring k by len(words[0]) -> map
func findSubstring(s string, words []string) []int {
	right := make(map[string]int, len(words))
	for _, word := range words {
		right[word] = right[word] + 1
	}
	unit := len(words[0])
	k := unit * len(words)
	ans := make([]int, 0, len(s))
	for i := 0; i <= len(s)-k; i += 1 {
		if valid(s[i:i+k], unit, right) {
			ans = append(ans, i)
		}
	}

	return ans
}

func valid(s string, unit int, words map[string]int) bool {
	left := make(map[string]int, len(words))
	for i := 0; i <= len(s)-unit; i += unit {
		sub := s[i : i+unit]
		left[sub] += 1
	}
	return EqualsMap(words, left)
}
