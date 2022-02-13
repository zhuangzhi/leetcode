package week6

// https://leetcode-cn.com/problems/word-break/
// 139. 单词拆分
func wordBreak(s string, wordDict []string) bool {
	dict := map[string]bool{}
	for _, w := range wordDict {
		dict[w] = true
	}
	dp := make([]bool, len(s)+1)
	// dp[i]代表s[0...i]能否被空格拆分成workDict中的词
	// dp[j] && dict[s[j:i]]，代表前s[0...j]能被拆分，并且s[j:i]也是在dict中即，
	// 在之前拆分的结果（到j）上又有一个单词s[j:i]
	dp[0] = true // 0长度单词是成功的
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
