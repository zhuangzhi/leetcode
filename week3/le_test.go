package week3

import (
	"fmt"
	"testing"

	. "github.com/zhuangzhi/leetcode/util"

	"github.com/stretchr/testify/assert"
)

func TestLe22(t *testing.T) {
	units := []struct {
		n      int
		Expect []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{4, []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}},
		{5, []string{"((((()))))", "(((()())))", "(((())()))", "(((()))())", "(((())))()", "((()(())))", "((()()()))", "((()())())", "((()()))()", "((())(()))", "((())()())", "((())())()", "((()))(())", "((()))()()", "(()((())))", "(()(()()))", "(()(())())", "(()(()))()", "(()()(()))", "(()()()())", "(()()())()", "(()())(())", "(()())()()", "(())((()))", "(())(()())", "(())(())()", "(())()(())", "(())()()()", "()(((())))", "()((()()))", "()((())())", "()((()))()", "()(()(()))", "()(()()())", "()(()())()", "()(())(())", "()(())()()", "()()((()))", "()()(()())", "()()(())()", "()()()(())", "()()()()()"}},
	}
	for _, u := range units {
		assert.ElementsMatch(t, u.Expect, generateParenthesis(u.n))
	}
}

func TestLe23(t *testing.T) {
	units := []struct {
		lists  [][]int
		Expect []int
	}{
		{[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{[][]int{}, []int{}},
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

func TestLe46(t *testing.T) {
	units := []struct {
		nums     []int
		expected [][]int
	}{
		// {[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{[]int{1, 2, 3}, [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}},
	}

	for _, u := range units {
		actual := permute(u.nums)
		assert.ElementsMatch(t, u.expected, actual)
	}
}

func TestLe47(t *testing.T) {
	units := []struct {
		nums     []int
		expected [][]int
	}{
		// {[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{[]int{1, 2, 3}, [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}}, {[]int{1, 1, 2}, [][]int{
			{1, 2, 1},
			{1, 1, 2},
			{2, 1, 1},
		}},
	}

	for _, u := range units {
		actual := permuteUnique(u.nums)
		fmt.Println(actual)
		fmt.Println(u.expected)
		assert.ElementsMatch(t, u.expected, actual)
	}
}

func TestLe77(t *testing.T) {
	units := []struct {
		n, k     int
		expected [][]int
	}{
		// {[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{4, 2, [][]int{
			{2, 4},
			{3, 4},
			{2, 3},
			{1, 2},
			{1, 3},
			{1, 4},
		}},
	}

	for _, u := range units {
		actual := combine(u.n, u.k)
		SortIntSlice(u.expected)
		SortIntSlice(actual)
		fmt.Println(u.expected)
		fmt.Println(actual)
		assert.ElementsMatch(t, u.expected, actual)
	}
}

func TestLe78(t *testing.T) {
	units := []struct {
		nums     []int
		expected [][]int
	}{
		// {[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{[]int{9, 0, 3, 5, 7}, [][]int{{}, {9}, {0}, {0, 9}, {3}, {3, 9}, {0, 3}, {0, 3, 9}, {5}, {5, 9}, {0, 5}, {0, 5, 9}, {3, 5}, {3, 5, 9}, {0, 3, 5}, {0, 3, 5, 9}, {7}, {7, 9}, {0, 7}, {0, 7, 9}, {3, 7}, {3, 7, 9}, {0, 3, 7}, {0, 3, 7, 9}, {5, 7}, {5, 7, 9}, {0, 5, 7}, {0, 5, 7, 9}, {3, 5, 7}, {3, 5, 7, 9}, {0, 3, 5, 7}, {0, 3, 5, 7, 9}}},
	}

	for _, u := range units {
		actual := subsets(u.nums)
		SortIntSlice(u.expected)
		SortIntSlice(actual)
		assert.ElementsMatch(t, u.expected, actual)
	}
}

func TestLe105(t *testing.T) {
	units := []struct {
		preorder, inorder []int
		expected          []int
	}{
		// {[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, []int{}},
	}

	for _, u := range units {
		root := buildTree(u.preorder, u.inorder)
		assert.NotNil(t, root)
		// assert.ElementsMatch(t, u.expected, actual)
	}
}
