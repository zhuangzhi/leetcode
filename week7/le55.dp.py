'''
https://leetcode-cn.com/problems/jump-game/
55. 跳跃游戏
'''
from typing import List
class Solution:
    def canJump1(self, nums: List[int]) -> bool:
        '''
        dp: 超时
        f[i]: 能否调到第i个位置
        f[i] = f[i] or (f[j] and j + nums[j]>=i)
        '''
        n = len(nums)
        f = [False] * (n)
        f[0] = True
        for i in range(1, n):
            can = False
            j = 0
            while j<i and not can:
                can = f[j] and j+nums[j] >= i
                j += 1
            f[i] = can
        return f[n-1]

    def canJump(self, nums: List[int]) -> bool:
        '''
        贪心
        '''
        n = len(nums)
        can = True
        maxn = nums[0]
        i = 1
        while can and i<n:
            can = maxn >= i
            maxn = max(i + nums[i], maxn)
            i+=1
        return can

s = Solution()
print(s.canJump([2,3,1,1,4]))
print(s.canJump([3,2,1,0,4]))
print(s.canJump([2,0,0]))