# week6，贪心，搜索，动态规划

贪心算法最快但不总是正确，先考虑搜索和动态规划了解整个状态空间的情况再考虑贪心，而且必须证明正确性。

## 贪心

每一步局部最优达到全局最优。

适用于每一步是最优解，全局也是最优解，所以需证明，有时候并不适用。如：
找零钱：如果零钱是：20,10,5可以使用贪心，但是如果是10,9,1就不适合。

数学归纳法，反证法。包容性，多判断一个，邻项交换法

### 包容性

如果零钱是：20,10,5
10的解可以🈶️多个5的解组成但不是最优，证明优先使用10更好

### 多判断一个

le45: 跳跃游戏 II, 本次跳跃和下次跳跃两次跳跃最远的为最优解。
122： 买卖股票的最佳时机 II，看明天是涨还是跌，涨就卖，跌就不买。

### 邻项交换法

求一个顺序是不是最优的。
如何证明：证明在这个序列中任意位置如果交换就会破坏最优解。如冒泡排序。
1665. 完成所有任务的最少初始能量：
必须有一个完整的证明：
// actual[p] - mini[p] <  actual[q] - mini[q]
// 任意位置p,q需要的能量 max(max(minimum[q], S+actual[q])+actual[p], minimum[p])
//  先q后p：max(max(minimum[p], S+actual[p])+actual[q], minimum[q])
// max(max(minimum[q], S+actual[p])+actual[p], minimum[p]) =>
// p,q: max(minimum[q] + actual[p], S+actual[p]+actual[q], minimum[p])
// q,p: max(minimum[p] + actual[q], S+actual[p]+actual[q], minimum[q])
// 比较上下两个大小，S+actual[p]+actual[q] 去掉
// p,q: max(minimum[q] + actual[p], minimum[q])
// q,p: max(minimum[p] + actual[q], minimum[q])
// 则若p,q较好：max(minimum[q] + actual[p], minimum[p])<max(minimum[p] + actual[q], minimum[q])
// 因为minimum[p] + actual[q] > minimum[p]， 所以这两项去掉
// 则若p,q较好：minimum[q] + actual[p] < minimum[p] + actual[q]
//  即：actual[p]- minimum[p] <  actual[q] - minimum[q]
// 因此可以按照 actual[i]- minimum[i]进行排序

## 动态规划

递归+记忆化搜索
三要素：阶段，状态，决策

```go
opt[0] = 0 //
// i, 阶段：线性增长
for i := 1; i <= amount; i++ {
    opt[i] = INF
    for j := 0; j < len(coins); j++ {
        if i-coins[j] >= 0 { // 这枚硬币满足要求
            // coins[j]决策，找到子问题
            // opt[i] 状态，具有最优子结构
            opt[i] = minInt(opt[i], opt[i-coins[j]]+1)
        }
    }
}
```
标准题解：
设opt[i]表示凑成金额i所需的最小硬币数
状态转移方程：  opt[i] = min(coin->coins){opt[i-coin]}
边界：opt[0] = 0, opt[i] = +INF (i>0)
目标：opt[amount]
时间复杂度：O(amount*coins)