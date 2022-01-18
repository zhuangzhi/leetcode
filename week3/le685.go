package week3

import (
	. "github.com/zhuangzhi/leetcode/util"
)

func findRedundantDirectedConnection(edges [][]int) []int {
	// result := make([]int, 0, len(edges))
	n := MaxEdgeNumber(edges)
	to := make([][]int, n+1)
	in := make([][]int, n+1)

	for _, edge := range edges {
		x, y := edge[0], edge[1]
		to[x] = append(to[x], y) //出度
		in[y] = append(in[y], x) //入度
	}
	root := -1
	var alt []int
	// 入度为0的点是根结点
	for idx, v := range in {
		if len(v) == 0 && idx != 0 {
			root = idx
		}
		if len(v) > 1 {
			// 入度大于1
			f := -1
			for _, k := range v {
				f = MaxInt(f, k)
			}
			alt = []int{f, idx}
		}
	}
	result := []int{}
	visited := make([]bool, n+1)
	var dfs func(x int) bool
	dfs = func(x int) bool {
		visited[x] = true
		for _, y := range to[x] {
			if visited[y] {
				result = []int{x, y}
				return true
			}
			if dfs(y) {
				return true
			}
		}
		visited[x] = false
		return false
	}
	if root < 0 {
		root = 1
	}
	dfs(root)
	if len(result) == 0 {
		return alt
	}
	return result
}
