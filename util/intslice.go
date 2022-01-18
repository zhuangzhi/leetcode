package util

import (
	"fmt"
	"sort"
)

func LessIntSlice(l, r []int) bool {
	return Ints(l).Less(r)
}

func SortIntSlice(intSlices [][]int) {
	IntsArray(intSlices).Sort()
}

type Int int

func (in Int) Compare(with int) int {
	switch {
	case int(in) < with:
		return -1
	case int(in) > with:
		return -1
	}
	return 0
}

func SortStrSlice(strslices [][]string) {
	for _, slice := range strslices {
		sort.Strings(slice)
	}
}

type Ints []int

func IntsFromRange(from, to int) Ints {
	s := make([]int, 0, to-from)
	for ; from < to; from++ {
		s = append(s, from)
	}
	fmt.Println(from, to, s)
	return s
}

func (is Ints) Len() int {
	return len(is)
}

func (is Ints) Copy() Ints {
	return append([]int(nil), is...)
}

func (is Ints) Append(vals ...int) Ints {
	is = append(is, vals...)
	return is
}

func (is Ints) RemoveEnd(n int) Ints {
	return is[:is.Len()-n]
}
func (is Ints) Pop() Ints {
	return is[:is.Len()-1]
}

func (is Ints) Sort() Ints {
	sort.Ints(is)
	return is
}

func (is Ints) ToItems() Items {
	items := NewItems(is.Len())
	is.ForEach(func(idx, val int) {
		items.Append(val, idx)
	})
	return items
}

func (is Ints) Slice() []int {
	return is
}

func (is Ints) ForEach(fn func(idx, val int)) {
	for idx, val := range is {
		fn(idx, val)
	}
}

func (is Ints) ComputeEach(fn func(idx, val int) int) Ints {
	for idx, val := range is {
		is[idx] = fn(idx, val)
	}
	return is
}

func (is Ints) ForEachBreakable(fn func(idx, val int) error) error {
	for idx, val := range is {
		if err := fn(idx, val); err != nil {
			return err
		}
	}
	return nil
}

func (is Ints) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is Ints) Equal(with []int) bool {
	return is.Compare(with) == 0
}

func (is Ints) ItemEqualWithPre(i int) bool {
	return i > 0 && is[i] == is[i-1]
}

func (is Ints) ReversLoopUntil(from int, fn func(idxR, val int) bool) (idx, val int) {
	for idx = from; idx < len(is) && idx >= 0; idx-- {
		val = is[idx]
		if !fn(idx, val) {
			return
		}
	}
	return
}

func (is Ints) LoopUntil(from int, fn func(idx, val int) bool) (idx, val int) {
	for idx = from; idx < len(is) && idx >= 0; idx++ {
		val = is[idx]
		if !fn(idx, val) {
			return
		}
	}
	return
}

func (is Ints) ElementMatch(with []int) bool {
	return is.ContainsAll(with) && Ints(with).ContainsAll(with)
}

func (is Ints) ContainsAll(with []int) bool {
	x := is.ValIdxMap()
	for _, v := range with {
		if _, ok := x[v]; !ok {
			return false
		}
	}
	return true
}

func (is Ints) ValIdxMap() map[int]int {
	x := map[int]int{}
	for idx, v := range is {
		x[v] = idx
	}
	return x
}

func (is Ints) Less(with []int) bool {
	return is.Compare(with) < 0
}

func (is Ints) Compare(with []int) int {
	l1, l2 := len(is), len(with)
	l := MinInt(l1, l2)
	for idx := 0; idx < l; idx++ {
		if v := Int(is[idx]).Compare(with[idx]); v != 0 {
			return v
		}
	}

	return Int(l1).Compare(l2)
}

func (is Ints) Sum() int {
	x := 0
	for _, v := range is {
		x += v
	}
	return x
}

type IntsArray [][]int

func (ia IntsArray) EmptyCopy() IntsArray {
	return make([][]int, 0, len(ia))
}

func (ia IntsArray) Copy() IntsArray {
	return append([][]int(nil), ia...)
}

func (ia IntsArray) Append(x ...[]int) IntsArray {
	return append(ia, x...)
}

func (ia IntsArray) RemoveEnd(n int) IntsArray {
	return ia[:len(ia)-n]
}

func (ia IntsArray) ForEach(fn func(idx int, s Ints)) {
	for idx, val := range ia {
		fn(idx, val)
	}
}

func (ia IntsArray) ComputeEach(fn func(i int, x Ints) Ints) IntsArray {
	for idx, val := range ia {
		ia[idx] = fn(idx, val)
	}
	return ia
}

func (ia IntsArray) Equal(with IntsArray) bool {
	ia.Sort()
	if len(ia) == len(with) {
		for idx, s := range ia {
			if !Ints(s).Equal(with[idx]) {
				return false
			}
		}
		return true
	}
	return false
}

func (ia IntsArray) SortItems() IntsArray {
	ia.ForEach(func(i int, x Ints) {
		x.Sort()
	})
	return ia
}

func (ia IntsArray) Sort() IntsArray {
	ia.SortItems()
	sort.Slice(ia, func(i, j int) bool {
		return Ints(ia[i]).Less(ia[j])
	})
	return ia
}
