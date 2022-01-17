"""
python 仅用于练习这个语言，没经过测试
"""
class Solution:
    def twoSum(self, nums, target):
        """
        hashmap 记录preval 对应的index，查看是不是有一个preval和当前值的和是target，
        如果存在就返回preval index 和当前index组成的数组
        """
        hashNum = dict()
        for i, v in enumerate(nums):
            need = target - v
            if need in hashNum: # 是否存在一个preval满足要求
                return [hashNum[need], i] # 返回这个preval的index及当前index组成的数组
            hashNum[v] = i
        return []
    def twoSumDoublePointer(self, nums, target):
        """
        双指针法
        """
        pairs = []
        for i, v in enumerate(nums):
            pairs.append((v,i))
        pairs.sort(key=lambda e:e[0])
        n = len(nums)
        second = n-1
        for first in range(n): # for in 避免 first++
            while first<second  and pairs[first][0] + pairs[second][0] < target:
                second-=1
            if first<second and pairs[first][0] + pairs[second][0] == target:
                return [pairs[first][1], pairs[second][1]]
        return []
