package lc_900

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 892. 三维形体的表面积

// 一个个计算吧，

func surfaceArea(grid [][]int) int {

	rs := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}

			rs += (grid[i][j]*4 + 2)

			// 上
			if i-1 >= 0 {
				rs -= min(grid[i][j], grid[i-1][j])
			}
			// 下
			if i+1 < len(grid) {
				rs -= min(grid[i][j], grid[i+1][j])
			}
			// 左
			if j-1 >= 0 {
				rs -= min(grid[i][j], grid[i][j-1])
			}
			// 右
			if j+1 < len(grid[0]) {
				rs -= min(grid[i][j], grid[i][j+1])
			}

		}
	}
	return rs

}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func TestSurfaceArea(t *testing.T) {

	assertions := assert.New(t)
	assertions.Equal(surfaceArea([][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}), 46)
	assertions.Equal(surfaceArea([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}), 32)
	assertions.Equal(surfaceArea([][]int{{1, 0}, {0, 2}}), 16)
	assertions.Equal(surfaceArea([][]int{{1, 2}, {3, 4}}), 34)
	assertions.Equal(surfaceArea([][]int{{2}}), 10)

}
