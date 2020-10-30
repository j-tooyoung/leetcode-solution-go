package main

// 岛屿的周长
// https://leetcode-cn.com/problems/island-perimeter/

// bruteforce 从右边和下边查看是否有相邻的1，如有
func islandPerimeter(grid [][]int) int {
	res := 0
	for i, col := range grid {
		for j, val := range col {
			if val == 1 {
				res += 4
				if j+1 < len(col) && grid[i][j+1] == 1 {
					res -= 2
				}
				if i+1 < len(grid) && grid[i+1][j] == 1 {
					res -= 2
				}
			}
		}
	}
	return res
}

// dfs
func islandPerimeter2(grid [][]int) int {
	for i, col := range grid {
		for j, val := range col {
			if val == 1 {
				return dfs(grid,i, j)
			}
		}
	}
	return 0
}

//
func dfs(grid [][]int, i int, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 1
	}
	if grid[i][j] == 2 {
		return 0
	}
	if grid[i][j] == 0 {
		return 1
	}
	grid[i][j] = 2
	return dfs(grid, i + 1, j) + dfs(grid, i - 1,j) + dfs(grid, i, j + 1) + dfs(grid, i, j - 1)
}

type pair struct{ x, y int }
var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func islandPerimeter3(grid [][]int) (ans int) {
	n, m := len(grid), len(grid[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
			ans++
			return
		}
		if grid[x][y] == 2 {
			return
		}
		grid[x][y] = 2
		for _, d := range dir4 {
			dfs(x+d.x, y+d.y)
		}
	}
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				dfs(i, j)
			}
		}
	}
	return
}
