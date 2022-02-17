from typing import List
import sys

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        inf = int(1e9)
        minprice, maxprofit = inf, 0
        for v in prices:
            maxprofit = max(v-minprice, maxprofit)
            minprice = min(v, minprice)
        return maxprofit

s = Solution()
r = s.maxProfit([7,1,5,3,6,4])
print(r)