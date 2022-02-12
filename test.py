from itertools import count


class Solution(object):
    def fourSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[List[int]]
        """
        ans = dict()
        nums.sort()
        l, r = 0, len(nums)-1
        while l < r-2:
            while l<r-2:
                l1 = l+1
                r1 = r-1
                while l1 < r1:
                    while l1 < r1:
                        v = nums[l] + nums[l1] + nums[r1] + nums[r]
                        if v == target:
                            xx = [nums[l], nums[l1], nums[r1], nums[r]]
                            ans[",".join([str(x) for x in xx])] = xx
                            r1-=1
                        elif v>target:
                            r1-=1
                        else:
                            break
                    l1+=1
                r-=1
            l+=1
            r = len(nums)-1
        out = []
        for v in ans.values():
            out.append(v)
        return out
s = Solution()
v = s.fourSum([1,0,-1,0,-2,2], 0)
print(v)
v = s.fourSum([2,2,2,2,2], 8)
print(v)
