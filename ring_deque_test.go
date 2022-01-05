package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRing(t *testing.T) {
	a := assert.New(t)
	que := NewRing(2)
	a.Equal(0, que.size)
	que.InsertFront(1)
	que.InsertLast(2)
	a.Equal(2, que.size)
	a.Equal(1, que.GetFront())
	que.DeleteFront()
	a.Equal(1, que.size)
	a.Equal(2, que.GetFront())
}
