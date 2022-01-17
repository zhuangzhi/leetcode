package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntsSwap(t *testing.T) {
	s1 := Ints([]int{1, 2, 3})
	s1.Swap(0, 1)
	assert.Equal(t, []int{2, 1, 3}, s1.Slice())
	s1.Sort()
	assert.Equal(t, []int{1, 2, 3}, s1.Slice())
}
