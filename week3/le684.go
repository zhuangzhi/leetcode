package week3

import . "github.com/zhuangzhi/leetcode/util"

func MaxEdgeNumber(edges [][]int) (n int) {
	n = 0
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		n = MaxInt(n, MaxInt(x, y))
	}
	return
}

func BuildUndirectedSideVector(maxEdge int, edges [][]int) (to [][]int) {
	to = make([][]int, maxEdge+1)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		to[x] = append(to[x], y) // 无向图，需要双向路径
		to[y] = append(to[y], x) // 无向图，需要双向路径
	}
	return
}

func findRedundantConnectionII(edges [][]int) []int {
	n := MaxEdgeNumber(edges)
	to := BuildUndirectedSideVector(n, edges)
	visited := make([]bool, n+1)
	path := make([]int, 0, n)
	var dfs func(x, father int) bool
	dfs = func(x, father int) bool {
		visited[x] = true
		for _, y := range to[x] {
			// 因为是无向图，双向指向的，忽略回指的路径。
			if y == father {
				continue
			}
			path = append(path, y)
			if !visited[y] {
				dfs(y, x)
			} else {
				return true
			}
		}
		visited[x] = false
		return false
	}

	return []int{}
}

func findRedundantConnection(edges [][]int) []int {
	n := MaxEdgeNumber(edges)
	to := make([][]int, n+1)
	visited := make([]bool, n+1)
	hasCycle := false
	var dfs func(x, father int)
	dfs = func(x, father int) {
		visited[x] = true
		for _, y := range to[x] {
			// 因为是无向图，双向指向的，忽略回指的路径。
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
