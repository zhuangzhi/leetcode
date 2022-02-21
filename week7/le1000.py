'''
https://leetcode-cn.com/problems/minimum-cost-to-merge-stones/
1000. 合并石头的最低成本
'''
class Solution:
    def mergeStones(self, stones: List[int], k: int) -> int:
        '''
        f[l, r, i] 把l->r 合并成 i堆的成本 
        f[l, r, 1] = f[l, r, k] + E(p:l->r)(nums[p])
        子问题 l, p, r l~p 合成 j 堆， p + 1 ~ r 合成 i-j堆，一共 i 堆
        f[l, r, i] = min[l<=p < r, 1<=j<i]{f[l,p,j] + f[p+1, r, i-j]}
        继续优化： j在什么位置没关系
        f[l, r, i] = min[p->[l,r)]{f[l, p, 1]+f[l, p, i-1]} (i>1)
        '''
        n = len(stones)
        # f k+1 * n * n 因为可以=K
        f =[[[1e9]* (k+1) for _ in range(n)] for _ in range(n)]
        
        for l in range(0, n):
            f[l][l][1] = 0
        
        for length in range(2, n+1): # 区间长度 2->n堆，最少要合并两堆，最多合并n堆
            for l in range(0, n-length+1): # 区间左端点：0 -> n - length
                r = l + length - 1         # 区间右端点： l + length
                for i in range(2, k+1):    # [l, r) 合成 i 堆的最小，i -> [2,K]
                    for p in range(l, r):  # 子问题 [l, p] [p]
                        f[l][r][i] = min(f[l][r][i], f[l][p][1] + f[p+1][r][i-1])
                sumn = sum(stones[l:r+1])
                f[l][r][1] = min(f[l][r][1], f[l][r][k] + sumn)

        return (f[0][n-1][1], -1)[f[0][n-1][1]==1e9]
