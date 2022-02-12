# week5

## 二分查找


## 不等式二分法

### 后继型

查找数列中>=target的数字，查找条件单调性：0,0,0,0,1,1,1，即前边不满足条件，周边的部分满足条件

#### 后继型程序模版

```go
// 后继型 用n表示无解
left, right := 0, len(nums)
for left < right {
    mid := (left + right) / 2
    if nums[mid] >= target {
        right = mid //后继型，一侧包含
    } else {
        left = mid + 1 // 一侧不包含
    }
}
```

### 前驱型

查找数列中<=target的数字，查找条件单调性：1,1,1,0,0,0,0,，即前边满足条件，周边的部分不满足条件

#### 前驱型程序模版

```go
// 前驱型 用-1表示无解
left, right = -1, len(nums)-1
for left < right {
    mid := (left + right + 1) / 2
    // 为什么+1向上取整
    // 当有[left,right]相邻的时候 (left+right)/2 =left，如果只有分支1
    // 即nums[mid] <= target，陷入死循环
    if nums[mid] <= target {
        left = mid // 前驱型，一侧包含
    } else {
        right = mid - 1 // 一侧不包含
    }
}
```

## 三分查找

单峰函数求极值

`mid`需要判断一下区间，是上升区间还是下降区间，比较`mid`和`mid+1`，如果是上升则极值不会在左侧，left=mid+1；否则极值不会在右侧，right=mid-1。

```go
left, right := 0, len(nums)-1
for left < right {
    lmid := (left + right) / 2
    rmid := lmid + 1
    if nums[lmid] <= nums[rmid] { // 上升区间，或两侧
        left = lmid + 1
    } else { // 下降区间，或两侧
        right = rmid - 1
    }
}
return right
```

## 排序

### 堆排序

```go

func heapSort(arr []int) {
    heap := NewHeapSlice(arr)
    for i:=0;heap.Size()>0;i++{
        arr[i] = heap.Top().(int)
        heap.Pop()
    }
}

```

### 归并排序

是一种基于分治的算法，时间复杂度O(NlogN)

```go

func mergeSort(arr []int, l, r int) {
    if l>=r {
        return
    }
    mid := (l+r)>>1
    mergeSort(arr, l, mid)
    mergeSort(arr, mid+1, r)
    merge(arr, l, mid, r)
}

func merge(arr []int, l, mid, r int) {
    temp := make([]int, r-l+1)
    i, j := l, mid+1
    for k := range temp {
        if j>r || (i<=mid && arr[i] <=arr[j]) {
            temp[k] = arr[i]
            i++
        } else {
            temp[k] = arr[j]
            j++
        }
    }
    for k:= range temp {
        arr[l+k] = temp[k]
    }
}

```

### 快速排序

* 选取中轴元素pivot
* 将小元素放到左边，大的元素放到右边
* 然后分别对左右进行快排

```go

func quickSort(arr []int, l, r int) {
    if l == r {
        return
    }
    pivot := partition(arr, l, r)
    quickSort(arr, l, pivot)
    quickSort(arr, pivot+1, r)
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

func partition(arr []int, l, r int) int {
    pivot := randInt(l, r)
    pivotVal := arr[pivot]
    for l <= r {
        for ; arr[l] < pivotVal; l++ {
        }
        for ; arr[r] > pivotVal; r-- {
        }
        if l == r {
            break
        }
        if l < r {
            arr[l], arr[r] = arr[r], arr[l]
            l++
            r--
        }
    }
    return r
}

```

### 计数排序 (Counting Sort)

### 桶排序 (Bucket Sort)

### 基数排序 （Radix Sort)

个位，十位，百位...顺序排序
