package week4

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	. "github.com/zhuangzhi/leetcode/util"
)

func TestLe23(t *testing.T) {
	units := []struct {
		lists  [][]int
		Expect []int
	}{
		{[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{[][]int{}, []int{}},
		{[][]int{[]int{}}, []int{}},
		{[][]int{[]int{}, []int{}}, []int{}},
	}
	for _, u := range units {
		lists := []*ListNode{}
		for _, l := range u.lists {
			lists = append(lists, BuildListNodes(l))
		}
		head := mergeKLists(lists)

		assert.Equal(t, u.Expect, StoreListNodes(head))
	}
}
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

func TestLe130(t *testing.T) {
	units := []struct {
		Grid   [][]byte
		Expect [][]byte
	}{
		{
			[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
			[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
		},
	}
	for _, u := range units {
		solve(u.Grid)
		assert.Equal(t, u.Expect, u.Grid)
	}
}

func TestLe329(t *testing.T) {
	// [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
	units := []struct {
		grid   [][]int
		Expect int
	}{
		{[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}, 4},
	}
	for _, u := range units {
		// assert.Equal(t, u.Expect, closedIsland(u.grid))
		assert.Equal(t, u.Expect, longestIncreasingPathDFS(u.grid))
	}
}

func TestLe355(t *testing.T) {
	twitter := Constructor()
	twitter.PostTweet(1, 5)
	fs := twitter.GetNewsFeed(1)
	assert.Equal(t, []int{5}, fs)
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	fs = twitter.GetNewsFeed(1)
	assert.Equal(t, []int{6, 5}, fs)
	twitter.Unfollow(1, 2)
	fs = twitter.GetNewsFeed(1)
	assert.Equal(t, []int{5}, fs)
}
func TestLe1254(t *testing.T) {
	// [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,0,1,0,1],[1,1,1,1,1,1,1,0]]
	units := []struct {
		grid   [][]int
		Expect int
	}{
		{[][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}, 2},
	}
	for _, u := range units {
		// assert.Equal(t, u.Expect, closedIsland(u.grid))
		assert.Equal(t, u.Expect, closedIslandII(u.grid))
	}
}
