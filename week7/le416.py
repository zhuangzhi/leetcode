from typing import List
"""
https://leetcode-cn.com/problems/partition-equal-subset-sum/
416. 分割等和子集
"""
class Solution:
    def canPartition(self, nums: List[int]) -> bool:
        """
        背包问题
        f[i][j]: 前i个数能选出一些数总和是j是否可行
        f[i,j] = f[i-1][j] or f[i-1][j-v[i]]
        """
        n = len(nums)
        s = sum(nums)
        print(s)
        if s%2==1:
            return False
        f = [False]*(s+1)
        f[0] = True
        print(f)
        for i in range(0, n): # [0, n)
            j = s//2
            while j>=nums[i]: #  [s/2-> num[i]]
                f[j] = f[j] or f[j-nums[i]]
                j -=1
                # print(f[j], j-nums[i], f[j-nums[i]])
        return f[s//2]

s = Solution()
v = s.canPartition([1,5,11,5])
print(v)
