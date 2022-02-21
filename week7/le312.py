'''
https://leetcode-cn.com/problems/burst-balloons/
312. 戳气球
'''
from typing import List
class Solution:
    def maxCoins(self, nums: List[int]) -> int:
        '''
        考虑最后戳哪个气球 p
        f[l, r] = max{f[l, p-1] + f[p+1, r] + nums[p]*nums[l-1]*nums[r+1]}
        以区间长度作为 dp “阶段” ， 区间端点作为dp “状态”
        计算n长度问题，先计算 l<n的问题
        '''
        n = len(nums)
        nums = [1] + nums +[1]
        f = [[0]*(n+2) for _ in range(n+2)]

        for len in range(1, n+1): # 最外层循环区间长度 1->n
            for l in range(n + 2 - len): # 循环区间左端点 0 -> n + 1 -len
                r = l + len - 1 #计算右侧端点
                for p in range(l, r+1): # 循环左右端点间的位置 l->r
                    f[l][r] = max(f[l][r], f[l][p-1] + f[p+1][r] + nums[p] * nums[l-1] * nums[r+1])
        return f[1][n]

    def maxCoins(self, nums: List[int]) -> int:
        '''
        考虑最后戳哪个气球 p
        f[l, r] = max{f[l, p-1] + f[p+1, r] + nums[p]*nums[l-1]*nums[r+1]}
        以区间长度作为 dp “阶段” ， 区间端点作为dp “状态”
        计算n长度问题，先计算 l<n的问题
        '''
        n = len(nums)
        s = [1] + nums + [1]
        f = [[-1]*(n+1) for _ in range(n+1)]

        def calc(l, r):
            if l>r:
                return 0
            if f[l][r] != -1:
                return f[l][r]
            for p in range(l, r+1): # l->r
                f[l][r] = max(f[l][r], calc(l, p-1) + calc(p+1, r) + s[p] * s[l-1] * s[r+1])
            return f[l][r]

        return calc(1, n)
