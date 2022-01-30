package lc_700

import (
	"fmt"
	"testing"
)

// 695. 岛屿的最大面积

// 深度优先搜索
func maxAreaOfIsland(grid [][]int) int {
	max := 0

	dx, dy := len(grid), len(grid[0])
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			if grid[x][y] == 1 {
				max = IntMax(max, dfs_maxAreaOfIsland(grid, x, y))
			}
		}
	}

	return max
}

func dfs_maxAreaOfIsland(grid [][]int, sx int, sy int) int {
	dx, dy := len(grid), len(grid[0])
	if sx < 0 || sx >= dx || sy < 0 || sy >= dy || grid[sx][sy] != 1 {
		return 0
	}
	count := 1
	// 设为0，避免重复扫描
	grid[sx][sy] = 0

	count += dfs_maxAreaOfIsland(grid, sx-1, sy)
	count += dfs_maxAreaOfIsland(grid, sx+1, sy)
	count += dfs_maxAreaOfIsland(grid, sx, sy-1)
	count += dfs_maxAreaOfIsland(grid, sx, sy+1)

	return count
}

func IntMax(i, j int) int {
	if i < j {
		return j
	} else {
		return i
	}
}

func Test_maxAreaOfIsland(t *testing.T) {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}

	fmt.Println(maxAreaOfIsland(grid))
}
