package week4

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	alphabet := make([]string, 10)
	alphabet[2] = "abc"
	alphabet[3] = "def"
	alphabet[4] = "ghi"
	alphabet[5] = "jkl"
	alphabet[6] = "mno"
	alphabet[7] = "pqrs"
	alphabet[8] = "tuv"
	alphabet[9] = "wxyz"
	ans := []string{}
	buf := make([]rune, 0, len(digits)+1)
	var dfs func(idx int, str []rune)
	dfs = func(idx int, str []rune) {
		if idx == len(digits) {
			ans = append(ans, string(append([]rune(nil), str...)))
			return
		}
		for _, c := range alphabet[digits[idx]-'0'] {
			dfs(idx+1, append(str, c))
		}
	}
	dfs(0, buf)
	return ans
}
