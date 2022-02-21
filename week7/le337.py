
'''
https://leetcode-cn.com/problems/house-robber-iii/
337. 打家劫舍 III
'''
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def rob(self, root: TreeNode) -> int:
        '''
        f[x, 0] 以x为根的子树，在不打劫x的情况下能盗取的最大金额
        f[x, 1] 以x为根的子树，在打劫x的情况下能盗取的最大金额
        '''
        f = {None:[0,0]}
        def dfs(node: TreeNode):
            if node==None:
                return
            f[node] = [0,0]
            dfs(node.left)
            dfs(node.right)
            left, right = f[node.left], f[node.right]
            f[node][0] = max(left[0], left[1]) + max(right[0], right[1])
            f[node][1] = left[0] + right[0] + node.val
        dfs(root)
        return max(f[root])