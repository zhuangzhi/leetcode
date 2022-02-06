package week4

import "fmt"

// https://leetcode-cn.com/problems/surrounded-regions/
// 130. 被围绕的区域

func solve(grid [][]byte) {
	row, col := len(grid), len(grid[0])
	var dfs func(x, y int, src, dst byte)
	directions := [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	dfs = func(x, y int, src, dst byte) {
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
		dfs(0, y, 'O', 'A')
		dfs(row-1, y, 'O', 'A')
	}
	// 把左右两边处理
	for x := range grid {
		dfs(x, 0, 'O', 'A')
		dfs(x, col-1, 'O', 'A')
	}
	fmt.Println(grid)
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 'O' {
				dfs(x, y, 'O', 'X')
			}
		}
	}
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 'A' {
				grid[x][y] = 'O'
			}
		}
	}

}
