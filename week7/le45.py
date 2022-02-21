from typing import List
class Solution:
    def jump1(self, nums: List[int]) -> int:
        n = len(nums)
        f = [1e9] * n
        f[0] = 0
        f[1] = 1
        for i in range(1, n):
            minn = 1e9
            for j in range(0, i):
                if i-j <= nums[j]:
                    minn = min(minn, f[j]+1)
            f[i] = minn
        return f[n-1]

    def jump(self, nums: List[int]) -> int:
        now = 0
        target = len(nums) - 1
        ans = 0
        while now < target:
            right = now + nums[now]
            if right >= target:
                return ans + 1
            next, nextRight = now, right
            for i in range(now+1, right+1):
                if i+nums[i] > nextRight:
                    nextRight = i + nums[i]
                    next = i
            now = next
            ans+=1

        return ans

s = Solution()
# print(s.jump([2,3,1,1,4]))
# print(s.jump([2,3,0,1,4]))
print(s.jump([1,2,3]))