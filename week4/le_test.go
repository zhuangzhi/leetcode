package week4

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLe51(t *testing.T) {
	units := []struct {
		n      int
		Expect string
	}{
		{1, "[[Q]]"},
		{4, "[[.Q.. ...Q Q... ..Q.] [..Q. Q... ...Q .Q..]]"},
	}
	for _, u := range units {
		actural := solveNQueens(u.n)
		assert.Equal(t, u.Expect, fmt.Sprint(actural))
	}
}
