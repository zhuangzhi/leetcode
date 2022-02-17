from typing import List

'''
https://leetcode-cn.com/problems/house-robber/
198. 打家劫舍
'''
class Solution:
    def rob(self, nums: List[int]) -> int:
        '''
        f[i][0]: 第i户没偷资产 = max(f[i-1][1],f[i-1][0])，前户偷了或没偷的资产较大的一个
        f[i][1]: 第i户偷了资产 = f[i-1][0] + nums[i-1]，前一户没偷的资产+i户的资产

        分两次做，第一次不偷第一个，第二次不偷第n个，取max
        '''
        n = len(nums)
        if n==1:
            return nums[0]
        s = nums
        f = [[-1e9,-1e9],[-1e9,-1e9]]
        f[1&1][0] = 0 
        for i in range (2, n+1):
            f[i&1][0] = max(f[i-1&1][1],f[i-1&1][0])
            f[i&1][1] = f[i-1&1][0] + s[i-1]
        ans = max(f[n&1][0], f[n&1][1])
        f = [[-1e9,-1e9],[-1e9,-1e9]]

        f[1&1][1] = nums[0]
        f[1&1][0] = 0 
        for i in range (2, n+1):
            f[i&1][0] = max(f[i-1&1][1],f[i-1&1][0])
            f[i&1][1] = f[i-1&1][0] + s[i-1]
        return max(ans, f[n&1][0])