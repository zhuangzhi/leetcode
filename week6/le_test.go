package week6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLe300(t *testing.T) {
	units := []struct {
		s        []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	}
	for _, u := range units {
		assert.Equal(t, u.expected, lengthOfLIS(u.s))
	}
}

func TestLe673(t *testing.T) {
	units := []struct {
		s        []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{1, 3, 5, 4, 7}, 2},
		{[]int{2, 2, 2, 2, 2}, 5},
	}
	for _, u := range units {
		assert.Equal(t, u.expected, findNumberOfLIS(u.s))
	}
}
