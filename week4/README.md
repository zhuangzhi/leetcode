# week4

## 地图相关

### 程序模版

```go

func serve(grid [][]int, expect int) int {
    row, col := len(grid), len(grid[0])
    // direction: 上下左右四个方向，
    directions := [][2]int{{0,1},{1,0},{0,-1},{-1,0}}

    var dfs func(x, y, expect, changeTo int)
    dfs =  func(x, y, expect, changeTo int) {
        // 从 x,y 分别向四个方向尝试
        for _, dir := range directions{
            nx := x + dir[0]
            ny := y + dir[1]
            if nx<0 || nx>=row || ny <0 || ny >= col || grid[nx][ny]!=expect {
                continue
            }
            grid[nx][ny] = changeTo
            dfs(nx, ny)
        }
    }
    ans := 0
    for x:=range grid {
        for y:=range grid[x] {
            if grid[x][y] == expect {
                dfs(x, y)
                ans++
            }
        }
    }
    return ans
}

```

## 二叉堆

二叉堆一般使用一个一位数组来存储，利用`完全二叉树`的节点编号特性

假设第一个元素存储下标为1的位置的话

* 索引为p的节点的左孩子的索引为p*2
* 索引为p的节点的右孩子的索引为p*2+1
* 索引为p的节点父亲的索引为p/2（下取整）

假设第一个元素存储下标为0的位置的话

* 索引为p的节点的左孩子的索引为p*2+1
* 索引为p的节点的右孩子的索引为p*2+2
* 索引为p的节点父亲的索引为(p-1)/2（下取整）

### 二叉堆插入

* 放在数组末尾
* 向上调整到了根停止 heapfiy up

```go
p -> p/2 -> p/4

```

### 删除（堆顶）

* heap[1] <-> heap[N]
* 向下调整到了叶子停止，选择左右两个中较大的一个交换

```go
// 类似 Sort 可以直接以函数形式存在

func HeapSlice(s interface{}, less func(i, j int) bool)
func HeapInts(s []int)

type heap struct {
    data []int
    min bool
}

func (h *heap) build(data []int) {

}

func (h *heap) insert(val int) {
    
}

func (h heap) root() {
    
}

func (h *heap) popRoot() {
    
}

func (h *heap) replaceRoot() {
    
}

```

## Priority Queue

可以二叉堆实现，或者+map
