package week4

// https://leetcode-cn.com/problems/number-of-islands/
// 200. 岛屿数量
// 网格/地图类搜索题目。
// 无向图
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	type pair struct {
		x, y int
	}
	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}
	visited := make([][]bool, m)
	for row := range grid {
		visited[row] = make([]bool, n)
	}

	var bfs func(row, col int)
	bfs = func(row, col int) {
		queue := []pair{}
		queue = append(queue, pair{row, col})
		visited[row][col] = true
		for len(queue) > 0 {
			temp := []pair{}
			for _, p := range queue {
				//
				for i := 0; i < 4; i++ { // 四个方向
					nx := p.x + dx[i]
					ny := p.y + dy[i]
					if nx < 0 || nx >= m || ny < 0 || ny >= n {
						continue
					}
					if grid[nx][ny] != '1' {
						continue
					}
					if visited[nx][ny] {
						continue
					}
					temp = append(temp, pair{nx, ny})
					visited[nx][ny] = true
				}
			}
			queue = temp
		}
	}

	ans := 0
	for idr, row := range grid {
		for idc, col := range row {
			if col == '1' && !visited[idr][idc] {
				//
				ans++
				bfs(idr, idc)
			}
		}
	}
	return ans
}

func numIslandsII(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	type pair struct {
		x, y int
	}
	dir := []pair{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	visited := make([][]bool, m)
	for row := range grid {
		visited[row] = make([]bool, n)
	}

	var dfs func(row, col int)
	dfs = func(row, col int) {
		for _, d := range dir {
			nx := d.x + row
			ny := d.y + col
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			if grid[nx][ny] != '1' {
				continue
			}
			if visited[nx][ny] {
				continue
			}
			visited[nx][ny] = true
			dfs(nx, ny)
		}
	}

	ans := 0
	for idr, row := range grid {
		for idc, col := range row {
			if col == '1' && !visited[idr][idc] {
				//
				ans++
				dfs(idr, idc)
			}
		}
	}
	return ans
}
