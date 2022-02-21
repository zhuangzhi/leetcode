'''
https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
124. 二叉树中的最大路径和
'''

'''
testcases
[1,2,3]
[-10,9,20,null,null,15,7]
[-3]
[-2,-1]
'''

from turtle import right
from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        '''
        f[node][0] - 不选择当前节点（已经构成完整路径a+b+c）
            case 1 左侧独立构成一个子序列: f[l][0]
            case 2 右侧独立构成一个子序列: f[r][0]
            case 3 左 + 右 + 中 构成子序列 其中 左右只在>0时选择 val + max(f[l][1], 0) + max(f[r][1], 0)
            max(case 1, case 2, case 3)
            f[node][0] = max(f[l][0], f[r][0],  val + max(f[l][1], 0) + max(f[r][1], 0))
        f[node][1] - 可以被选择max(b, c, 0) + a
        f[node][1] = max(f[l][1], f[r][1], 0) + val
        '''
        f = {None:[0,0]}
        def dfs(node: Optional[TreeNode]):
            if node==None:
                return
            f[node] = 0
            dfs(node.left)
            dfs(node.right)
            left, right = f[node.left], f[node.right]
            val = node.val
            gain = max(left[1], 0) + max(right[1], 0) + val
            f[node][0] = max(left[0], right[0], gain)
            f[node][1] = max(left[1], right[1], 0) + val # 左右选一个或不选
        dfs(root)
        return max(f[root])
    def maxPathSumII(self, root: Optional[TreeNode]) -> int:
        maxSum = [-1e9]
        def dfs(node: Optional[TreeNode]) -> int:
            if node==None:
                return 0
            leftGain = max(dfs(node.left), 0)
            rightGain = max(dfs(node.right), 0)
            val = node.val
            priceNewPath = val + leftGain + rightGain # a + b + c max
            maxSum[0] = max(maxSum[0], priceNewPath)
            
            return val + max(leftGain, rightGain)
        dfs(root)
        return maxSum[0]

