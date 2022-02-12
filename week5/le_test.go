package week5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLe327(t *testing.T) {
	units := []struct {
		nums         []int
		lower, upper int
		Expect       int
	}{
		{[]int{-2, 5, -1}, -2, 2, 3},
		// {[]int{0}, 0, 0, 1},
	}
	for _, u := range units {
		assert.Equal(t, u.Expect, countRangeSum(u.nums, u.lower, u.upper))
	}
}

func TestLe410(t *testing.T) {
	units := []struct {
		nums   []int
		m      int
		Expect int
	}{
		{[]int{7, 2, 5, 10, 8}, 2, 18},
	}
	for _, u := range units {
		assert.Equal(t, u.Expect, splitArray(u.nums, u.m))
	}
}
