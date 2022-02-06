package week4

import "fmt"

// 1254. 统计封闭岛屿的数目
// 先把边缘岛屿消除，在查找所有岛屿
func closedIsland(grid [][]int) int {
	ans := 0
	row, col := len(grid), len(grid[0])
	var dfs func(i, j, src, dst int)
	directions := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	dfs = func(x, y, src, dst int) {
		if grid[x][y] != src {
			return
		}
		grid[x][y] = dst

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || ny < 0 || nx >= row || ny >= col || grid[nx][ny] != src {
				continue
			}
			dfs(nx, ny, src, dst)
		}
	}
	// 把顶和底处理
	for y := range grid[0] {
		dfs(0, y, 0, 2)
		dfs(row-1, y, 0, 2)
	}
	// 把左右两边处理
	for x := range grid {
		dfs(x, 0, 0, 2)
		dfs(x, col-1, 0, 2)
	}

	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 0 {
				dfs(x, y, 0, 2)
				ans++
			}
		}
	}
	return ans
}

func closedIslandII(grid [][]int) int {
	fmt.Println(grid)
	ans := 0
	row, col := len(grid), len(grid[0])
	directions := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	var bfs func(i, j, src, dst int)
	bfs = func(x, y, src, dst int) {
		queue := [][2]int{{x, y}}

		for len(queue) > 0 {
			tmp := [][2]int{}
			for _, p := range queue {
				x, y = p[0], p[1]
				if grid[x][y] != src {
					return
				}
				grid[x][y] = dst

				for _, dir := range directions {
					nx, ny := x+dir[0], y+dir[1]
					if nx < 0 || ny < 0 || nx >= row || ny >= col || grid[nx][ny] != src {
						continue
					}
					tmp = append(tmp, [2]int{nx, ny})
				}
			}
			queue = tmp
		}
	}
	// 把顶和底处理
	for y := range grid[0] {
		bfs(0, y, 0, 2)
		bfs(row-1, y, 0, 2)
	}
	// 把左右两边处理
	for x := range grid {
		bfs(x, 0, 0, 2)
		bfs(x, col-1, 0, 2)
	}
	fmt.Println(grid)
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 0 {
				bfs(x, y, 0, 2)
				ans++
			}
		}
	}
	return ans
}
