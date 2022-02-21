'''
https://leetcode-cn.com/problems/coin-change-2/
518. 零钱兑换 II
'''
from typing import List
class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        '''
        f[i,j]: 第i个coins选择的情况下达到j金额方案
        f[i,j] = f[i-1][j] + f[i-1][j-coins[i]]
        可以多选，所以是正循环
        '''
        k = len(coins)
        f = [0] * (amount+1)
        f[0] = 1
        for i in range(0, k):
            for j in range(coins[j], amount+1):
                f[j] = f[j] + f[j-coins[i]]
        return f[amount]
