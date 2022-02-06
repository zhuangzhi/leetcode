package week4

// https://leetcode-cn.com/problems/count-vowels-permutation/
// 1220. 统计元音字母序列的数目
func countVowelPermutation(n int) int {
	array := []int{'a', 'e', 'i', 'o', 'u'}
	follow := [128][]int{}
	follow['a'] = []int{'e'}
	follow['e'] = []int{'a', 'i'}
	follow['i'] = []int{'a', 'e', 'o', 'u'}
	follow['o'] = []int{'i', 'u'}
	follow['u'] = []int{'a'}
	ans := 0
	// perm := make([]int, n)
	var dfs func(idx, pre int)
	dfs = func(idx, pre int) {
		candidate := array
		if idx > 0 {
			candidate = follow[pre]
		}
		if idx == n-1 {
			ans += len(candidate)
			return
		}
		for _, c := range candidate {
			dfs(idx+1, c)
		}
	}
	dfs(0, -1)
	return ans
}
