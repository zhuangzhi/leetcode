package util

import "sort"

type Item struct {
	First, Second int
}

type Items []Item

func NewItems(cap int) Items {
	return make([]Item, 0, cap)
}

func (items *Items) Append(first, second int) *Items {
	*items = append(*items, Item{First: first, Second: second})
	return items
}

func (items Items) Len() int {
	return len(items)
}

func (items *Items) RemoveLast() {
	*items = (*items)[:len(*items)-1]
}
func (items *Items) SortReverse() {
	sort.Slice(*items, func(i, j int) bool {
		return (*items)[i].First >= (*items)[j].First
	})
}

func (items *Items) Sort() {
	sort.Slice(*items, func(i, j int) bool {
		return (*items)[i].First < (*items)[j].First
	})
}
func (items *Items) SortBySecond() {
	sort.Slice(*items, func(i, j int) bool {
		return (*items)[i].Second < (*items)[j].Second
	})
}
