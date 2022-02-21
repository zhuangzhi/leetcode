
'''
https://leetcode-cn.com/problems/perfect-squares/
279. 完全平方数
'''

import math

class Solution:
    def numSquares(self, n: int) -> int:
        '''
        f[i] 需要最少多少个平方数和是i (0->n)
        j 自然数 0->sqrt(i) 是不是被选中
        '''
        f = [0]*(n+1)
        for i in range(1, n+1):
            j = 1
            minn = 1e9 # min = maxint
            while j*j <=i:
                minn= min(minn, f[i-j*j])
                j += 1
            f[i] = minn + 1
        return f[n]
s = Solution()
print(s.numSquares(112))