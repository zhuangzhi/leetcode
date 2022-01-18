package week3

import "github.com/zhuangzhi/leetcode/util"

func findRedundantConnection(edges [][]int) []int {
	n := 0
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		n = util.MaxInt(n, util.MaxInt(x, y))
	}
	to := make([][]int, n+1)
	visited := make([]bool, n+1)
	hasCycle := false
	var dfs func(x, father int)
	dfs = func(x, father int) {
		visited[x] = true
		for _, y := range to[x] {
			if y == father {
				continue
			}
			if !visited[y] {
				dfs(y, x)
			} else {
				hasCycle = true
			}
		}
		visited[x] = false
	}
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		to[x] = append(to[x], y)
		to[y] = append(to[y], x)
		dfs(x, 0)
		if hasCycle {
			return edge
		}
	}
	return []int{}
}
