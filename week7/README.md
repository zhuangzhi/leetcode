# 动态规划II

## 股票

121,122,123,188,714,309

i：天数，j：是否持仓（0，1），K：已交易次数，l：是否在冷冻期（0，1）

f(i, j, k, l)

### i, j

f(i, j)
决策：
* 买 f (i, 1) <- f(i-1, 0)-price[i]
* 卖 f (i, 0) <- f(i-1, 1)+price[i]
* 不动 f (i, 1) <- f(i-1, 0)

求max
f(0, 0) = -inf (-1e9)，（求min inf）
f (i) = max(f)

#### 空间优化

只有 i, i-1有用，使用滚动数组
第一步，每个i/i-1后边都加上&1

```go
f[i&1].. = f[(i-1)&1] ...
```

## 背包问题

f[i,j]表示从前i个物品中选取了总体积为j的物品放入背包，物品的最大总价值，所以i(0, n), j(0, m)
f[i,j] = max{f[i-1, j](不选择第i个),f[i-1, j-v[i]]+w[i](选择第i个)}

```python
f = [-1e9 for j in range(m)] for i in range(n)
f[0][0] = 0

for i in range(n):
    for j in range(v[i], m):
        f[i][j] = f[i-1][j]
    for j in range(v[i], m):
        f[i][j] = max(f[i][j], f[i-1][j-v[i]] + w[i])
return max(f[n-1])
```

### 背包问题优化

```python
f = [-1e9 for x in range(m+1)]
f[0] = 0

for i in range(n):
    j = m
    while j>=v[i]:
        f[j] = max(f[j], f[j-v[i]] + w[i])
        j -= 1
return max(f)
```

```go
func take(m int, v, w []int) {
    // m : 体积为m的背包
    // v := []int{...} // 第i个物品的体积
    // w := []int{...} // 第i个物品的价值
    n := len(v)
    f := make([]int, m+1)
    for i := range f{
        f[i] = -1e9
    }
    f[0] = 0
....
    
}

```
