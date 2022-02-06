package util

import (
	"fmt"
	"reflect"
)

type SliceHeap struct {
	slice reflect.Value
	swap  func(i, j int)
	less  func(i, j interface{}) bool
}

func (s *SliceHeap) Len() int {
	return s.slice.Len()
}

func (s *SliceHeap) Index(i int) interface{} {
	return s.slice.Index(i).Interface()
}

func (s *SliceHeap) Swap(i, j int) {
	s.swap(i, j)
}
func (s *SliceHeap) Slice() interface{} {
	return s.slice.Interface()
}

func (s *SliceHeap) Less(i, j int) bool {
	i0, j0 := s.slice.Index(i), s.slice.Index(j)
	return s.less(i0.Interface(), j0.Interface())
}

func (s *SliceHeap) Pop() interface{} {
	n := s.Len() - 1
	e := s.Index(n)
	s.slice = s.slice.Slice(0, n)
	s.swap = reflect.Swapper(s.slice.Interface())
	return e
}
func (s *SliceHeap) Push(x interface{}) {
	s.slice = reflect.Append(s.slice, reflect.ValueOf(x))
	s.swap = reflect.Swapper(s.slice.Interface())
}

func IterateSlice(v interface{}, fn func(interface{})) {
	s := reflect.ValueOf(v)
	n := s.Len()
	for i := 0; i < n; i++ {
		x := s.Index(i)
		fn(x.Interface())
	}
}

func PrintSlice(x interface{}) {
	IterateSlice(x, func(i interface{}) {
		fmt.Printf("(%T, %v)", i, i)
	})
	fmt.Println()
}

func HeapSlice(v interface{}, less func(i, j interface{}) bool) Heap {
	sh := &SliceHeap{
		slice: reflect.ValueOf(v),
		less:  less,
		swap:  reflect.Swapper(v),
	}
	heap := Heap{
		Less:     sh.Less,
		Len:      sh.Len,
		Swap:     sh.Swap,
		PopFunc:  sh.Pop,
		PushFunc: sh.Push,
		Index:    sh.Index,
		Slice:    sh.Slice,
	}
	heap.Init()
	return heap
}

// heap

type Heap struct {
	Swap     func(i, j int)
	Less     func(i, j int) bool
	Len      func() int
	PopFunc  func() interface{} // Get and remove last item.
	PushFunc func(x interface{})
	Index    func(i int) interface{}
	Slice    func() interface{}
}

func (heap Heap) Init() {
	n := heap.Len()
	for i := n/2 - 1; i >= 0; i-- {
		heap.down(i)
	}
}

func (heap Heap) Fix(i int) {
	if !heap.down(i) {
		heap.up(i)
	}
}

func (heap Heap) Empty() bool {
	return heap.Len() == 0
}

func (heap Heap) Top() interface{} {
	if heap.Empty() {
		return nil
	}
	return heap.Index(0)
}

func (heap Heap) ReplaceTop(x interface{}) interface{} {
	n := heap.Len() - 1
	if n < 0 {
		heap.Push(x)
		return nil
	}
	heap.Swap(0, n)
	r := heap.PopFunc()
	heap.PushFunc(x)
	heap.Swap(0, n)
	heap.down(0)
	return r
}

func (heap Heap) Pop() interface{} {
	n := heap.Len() - 1
	if n < 0 {
		return nil
	}
	if n > 0 {
		heap.Swap(0, n)
	}
	r := heap.PopFunc()
	if heap.Len() > 0 {
		heap.down(0)
	}
	return r
}

func (heap Heap) Remove(i int) interface{} {
	if i > 0 && i < heap.Len() {
		n := heap.Len() - 1
		if i != n {
			heap.Swap(i, n)
			r := heap.PopFunc()
			heap.down(i)
			return r
		}
		return heap.PopFunc()
	}
	return nil
}

func (heap Heap) Push(x interface{}) {
	heap.PushFunc(x)
	heap.up(heap.Len() - 1)
}

func (heap Heap) up(p int) {
	for {
		f := (p - 1) / 2
		if f == p || heap.Less(f, p) {
			break
		}
		heap.Swap(f, p)
		p = f
	}
}

func (heap Heap) down(p0 int) bool {
	i := p0
	for {
		j := i*2 + 1
		if j >= heap.Len() || j < 0 {
			break
		}
		if j1 := j + 1; j1 < heap.Len() && heap.Less(j1, j) {
			j = j1
		}
		if heap.Less(i, j) {
			break
		}
		heap.Swap(i, j)
		i = j
	}
	return i > p0
}
