package week4

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix/
// 329. 矩阵中的最长递增路径
// 输入：matrix = [[9,9,4],[6,6,8],[2,1,1]]
// 输出：4
// 解释：最长递增路径为 [1, 2, 6, 9]。
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 200
// 0 <= matrix[i][j] <= 231 - 1

// BFS->图，由顶至下，一层一层计算，合适最短路径

func num(x, y, n int) int {
	return x*n + y
}

func valid(x, y, m, n int) bool {
	return x >= 0 && y >= 0 && x < m && y < n
}

// DFS 记忆画搜索，如果路由到一个节点这个歌点之后可以有几个已经有记录，直接加上，停止搜索
func longestIncreasingPath(matrix [][]int) int {
	// 图排序
	m, n := len(matrix), len(matrix[0])
	directions := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	out := make([][]int, m*n)  // 出边
	degree := make([]int, m*n) // 入度
	dist := make([]int, m*n)
	queue := []int{}
	var addEdge = func(from, to int) {
		degree[to]++
		out[from] = append(out[from], to)
	}
	for x := range matrix {
		for y, val := range matrix[x] {
			for _, dir := range directions {
				nx, ny := x+dir[0], y+dir[1]
				if valid(nx, ny, m, n) && val < matrix[nx][ny] {
					addEdge(num(x, y, n), num(nx, ny, n))
				}
			}
		}
	}

	for i := 0; i < m*n; i++ {
		if degree[i] == 0 { // 入度为0， 是root
			queue = append(queue, i)
			dist[i] = 1
		}
	}

	for len(queue) > 0 {
		tmp := []int{}
		for _, x := range queue {
			for _, y := range out[x] {
				degree[y]-- //
				dist[y] = MaxInt(dist[y], dist[x]+1)
				if degree[y] == 0 { // 最晚的一个已经做完了
					tmp = append(tmp, y)
				}
			}
		}
		queue = tmp
	}
	ans := 0
	for _, v := range dist {
		ans = MaxInt(ans, v)
	}
	return ans
}

func longestIncreasingPathDFS(matrix [][]int) int {
	// DFS + 记忆化搜索
	m, n := len(matrix), len(matrix[0])
	directions := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	dist := make([][]int, m) // default 0, 没算过，否则至少是1（它本身）
	for row := range matrix {
		dist[row] = make([]int, n)
	}

	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if v := dist[x][y]; v != 0 {
			return v
		}
		dist[x][y] = 1 // 初始为1，他自己
		for _, d := range directions {
			nx, ny := x+d[0], y+d[1]
			if valid(nx, ny, m, n) && matrix[nx][ny] > matrix[x][y] {
				r := dfs(nx, ny) + 1
				dist[x][y] = MaxInt(dist[x][y], r)
			}
		}
		return dist[x][y]
	}
	ans := 0
	for x := range matrix {
		for y := range matrix[x] {
			ans = MaxInt(ans, dfs(x, y))
		}
	}
	return ans
}
