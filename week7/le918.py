'''
https://leetcode-cn.com/problems/maximum-sum-circular-subarray/
918. 环形子数组的最大和
'''
from typing import List
class Solution:
    def maxSubarraySumCircular2(self, nums: List[int]) -> int:
        '''
        '''
        n = len(nums)
        s = [0]*(2*n+1)
        for i, v in enumerate(nums):
            s[i+1] = s[i] + v
        for i, v in enumerate(nums):
            s[n+i+1] = s[n+i] + v
        ans = -1e9
        q = []
        for i in range(1, 2*n+1):
            while q and q[0] < i-n:
                q.pop(0)
            if q:
                ans = max(ans, s[i] - s[q[0]])
            while q and s[q[-1]] >= s[i]:
                q.pop()
            q.append(i)
        return ans

    def maxSubarraySumCircular(self, nums: List[int]) -> int:
        n = len(nums)
        s = [0]*(n+1)
        for i, v in enumerate(nums):
            s[i+1] = s[i] + v
        temp, ans = 1e9, -1e9
        for i in range(2, n+1):
            temp = min(temp, s[i-1])
            ans = max(ans, s[i] - temp)

        temp, ansMin = -1e9, 1e9
        for i in range(2, n+1):
            temp = max(temp, s[i-1])
            ansMin = min(ansMin, s[i] - temp)
        ans = max(ans, s[n]-ansMin)
        return ans
    
s = Solution()
v = s.maxSubarraySumCircular([1,-2,3,-2])
print(v)
v = s.maxSubarraySumCircular([5,-3,5])
print(v)
v = s.maxSubarraySumCircular([3,-2,2,-3])
print(v)
