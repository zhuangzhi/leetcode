package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	h := HeapSlice([]int{}, func(i, j interface{}) bool {
		return i.(int) < j.(int)
	})
	assert.Equal(t, 0, h.Len())
	assert.Nil(t, h.Pop())
}
func TestHeapPop(t *testing.T) {
	s := Range(0, 10)
	heap := HeapSlice(s, func(i, j interface{}) bool {
		return i.(int) < j.(int)
	})
	for _, expect := range Range(0, 10) {
		v := heap.Pop().(int)
		assert.Equal(t, expect, v)
	}
}
func TestHeapPush(t *testing.T) {
	s := []int{}
	heap := HeapSlice(s, func(i, j interface{}) bool {
		return i.(int) < j.(int)
	})
	for _, i := range Range(9, 0) {
		heap.Push(i)
		fmt.Println(heap.Slice())
	}
	for _, expect := range Range(1, 10) {
		v := heap.Pop().(int)
		assert.Equal(t, expect, v)
	}
	// assert.Nil(t, "")
}

func TestHeapReplace(t *testing.T) {
	s := []int{}
	heap := HeapSlice(s, func(i, j interface{}) bool {
		return i.(int) < j.(int)
	})
	for i := 9; i > 0; i-- {
		heap.Push(i)
		fmt.Println(heap.Slice())
	}
	for i := 10; i < 19; i++ {
		heap.ReplaceTop(i)
		assert.Equal(t, i-8, heap.Top())
		fmt.Println(heap.Slice())
	}
}
