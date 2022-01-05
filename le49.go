package leetcode

import "sort"

// https://leetcode-cn.com/problems/group-anagrams/
// 49. 字母异位词分组

func groupAnagrams(strs []string) [][]string {
	hash := map[string][]string{}
	for _, str := range strs {
		key := []byte(str)
		sort.Slice(key, func(i, j int) bool {
			return key[i] < key[j]
		})
		s := string(key)
		hash[s] = append(hash[s], str)
	}
	ret := make([][]string, 0, len(hash))
	for _, v := range hash {
		ret = append(ret, v)
	}
	return ret
}
