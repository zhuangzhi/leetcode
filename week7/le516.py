'''
https://leetcode-cn.com/problems/longest-palindromic-subsequence/
516. 最长回文子序列
'''

class Solution:
    def longestPalindromeSubseq(self, s: str) -> int:
        '''
        建立n*n的二维矩阵，f[i][j]代表 字符串s[i, j]最长回文子序列
        f[i][i] = 1
        如果s[i] == s[j]，则f[i][j] = f[i+1][j-1] + 2 即： x, x+1, .... j-1, j (x+1,j-1)两边各加一个。
        f[i][j] = max(f[i, j-1], f[i+1, j])
        '''
        n = len(s)
        f = [[0]*n for _ in range(n)]
        for i in range(n-1, -1, -1):
            f[i][i] = 1 # 自己和自己是回文
            for j in range(i+1, n):
                if s[i]==s[j]:
                    f[i][j] = f[i+1][j-1] + 2 # i, (i+1, j-1), j， 前一个基础上，左右各加一
                else:
                    f[i][j] = max(f[i+1][j], f[i][j-1]) # 选择偏向i的一边或j的一边
        return f[0][n-1]