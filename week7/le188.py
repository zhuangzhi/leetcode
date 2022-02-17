from typing import List
class Solution:
    def maxProfit(self, K: int, prices: List[int]) -> int:
        n = len(prices)
        # Move index to -1 based index
        prices = [0] + prices

        # 1. Define f, initial inf
      
        # Don't use [[[-1e9]*(K+1)]*2]*(n+1), Shallow copy
        # Use deep copy [[[-1e9]*(K+1)]*2]*(n+1)
        f = [[[-1e9 for k in range(0, K+1)] for j in range(0, 2)] for i in range(n+1)]
        f[0][0][0] = 0

        # 2, Loop over all states
        for i in range(1, n+1):
            price = prices[i]
            # Copy decisions
            for k in range(0, K+1):
                if k>0:
                    f[i][1][k] = max(f[i-1][1][k], f[i-1][0][k-1] - price)
                f[i][0][k] = max(f[i-1][0][k], f[i-1][1][k] + price)
        # 3. Return
        return max(f[n][0])

s = Solution()
print(s.maxProfit(2, [2,4,1]))