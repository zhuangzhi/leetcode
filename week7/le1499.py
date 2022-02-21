'''
https://leetcode-cn.com/problems/max-value-of-equation/
1499. 满足不等式的最大值
'''
from typing import List

class Solution:
    def findMaxValueOfEquation(self, points: List[List[int]], k: int) -> int:
        '''
        坐标 x 的值从小到大排序
        premin/max
        求 y[i] + y[j] +|x[i]-x[j]| 最大，且 |x[i]-x[j]| <=k
        i > j => y[i] + x[i] + y[j] - x[j]
        维护一个y[j] - x[j] 的递减队列，
        '''
        ans = -1e9
        
        '''
        queue : 上界：x[j] >= x[i]-k
                下界：j<=i-1
        '''
        q = []
        for i, p in enumerate(points):
            # 清除不符合条件的队列内容 x2-x1 > k两点间距离>k
            while q and points[q[0]][0] < p[0]-k:
                q.pop(0) # pop 左侧pop

            # ans = max(ans, y[i] + x[i] + y[j] - x[j])
            if q:
                x = points[q[0]]
                ans = max(ans, p[1] + p[0] + x[1]-x[0])
            
            '''
            维护queue单调性是 y[j]-x[j]的递减队列
            如果当前队列中y-x < 当前y-x就去除掉
            '''
            while q and \
                points[q[-1]][1] - points[q[-1]][0] <= p[1]-p[0]:
                q.pop() # 右侧pop y1-x1 < y-x 
            q.append(i) # 把当前点加入到队尾
        return ans
