package main

import (
	"fmt"
	"math"

	"github.com/zhuangzhi/leetcode"
)

const N = 1e5 + 10
const (
	Min = -3e9
	Max = 3e9
)

func acWing(a leetcode.Items) leetcode.Items {
	fmt.Println(a)
	a.Sort()
	fmt.Println(a)

	n := len(a) - 2
	l := [N]int{}
	r := [N]int{}
	p := [N]int{}

	for i := 1; i <= n; i++ {
		fmt.Println(a[i])
		l[i] = i - 1
		r[i] = i + 1       //记录左右相邻元素的位置
		p[a[i].Second] = i //记录排序后各元素位置
	}
	ans := leetcode.NewItems(n)
	for i := n; i > 1; i-- {
		j := p[i]
		left, right := l[j], r[j]
		leftV := math.Abs(float64(a[j].First - a[left].First))
		rightV := math.Abs(float64(a[j].First - a[right].First))
		if leftV < rightV {
			ans.Append(int(leftV), a[left].Second)
		} else {
			ans.Append(int(rightV), a[right].Second)
		}
		l[right] = left
		r[left] = right
	}
	return ans
}

func readInput() leetcode.Items {
	var n, a int
	fmt.Print("Enter n: ")
	fmt.Scanf("%d", &n)
	items := leetcode.NewItems(n + 2)
	items.Append(Min, 0)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &a)
		items.Append(a, i)
	}
	items.Append(Max, 0)
	return items
}

func main() {
	nums := readInput()

	ans := acWing(nums)
	// n := len(nums) - 2
	for _, l := range ans {
		fmt.Println(l.First, l.Second)
	}
}
