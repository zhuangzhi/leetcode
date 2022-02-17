from typing import List
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
       
        n = len(prices)
        # Move index to -1 based index
        prices = [0] + prices

        # 1. Define f, initial inf
        
        """
        f[i][0] = f[i-1][1] + prices[i] ss=> f[i-1][0] - price[i-1] + price[i] 昨天买了股票，今天赚了
                || f[i-1][0] 继续不持仓
        f[i][0]（今天没有持仓） = f[i-1][1] + prices[i] 昨天持仓，今天卖掉了的资产
                         或者 = f[i-1][0] 昨天没有持仓，今天也没交易
        就是比较 昨天买了今天卖了是亏还是赚，选择大的那个
        f[i][1] 即今天持仓 比较昨天卖了，今天再买回来的资产和昨天就持仓的情况，选择大的一个。
        """

        f = [[-1e9 for j in range(0, 2)] for i in range(n+1)]
        f [0][0] = 0

        # 2, Loop over all states
        for i in range(1, n+1):
            price = prices[i]
            # Copy decisions
            f[i][1] = max(f[i-1][1], f[i-1][0] - price)
            f[i][0] = max(f[i-1][0], f[i-1][1] + price)

        # 3. Return 
        return f[n][0]