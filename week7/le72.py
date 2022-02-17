from typing import List

'''
https://leetcode-cn.com/problems/edit-distance/
72. 编辑距离
'''
class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        '''
        f[i][j]: word1 前i个字符变成 work2前j个字符最少需要多少次操作
        f[i][j] = min(f[i][j-1]+1(插入), f[i-1][j]+1(删除), f[i-1][j-1](替换+1/不变+0):)
        初值：f[i][0] = i, f[0][j] = j
        目标：f[i][j]
        '''
        m, n = len(word1), len(word2)
        word1, word2 = " " + word1, " " + word2
        f = [[1e9 for j in range(n+1)] for i in range(m+1)]
        for i in range(m+1):
            f[i][0] = i
        for j in range(n+1):
            f[0][j] = j
        for i in range(1, m+1):
            eq = (1, 0)[word1[i] == word2[j]] # (false, true)[condition]
            for j in range(1, n+1):
                f[i][j] = min(f[i][j-1]+1, f[i-1][j]+1, f[i-1][j-1] + eq)

        return f[m][n]


