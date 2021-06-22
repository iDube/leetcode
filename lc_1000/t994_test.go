package lc_1000

import (
	"fmt"
	"math"
	"testing"
)

type orange struct {
	x, y int
}

// 想到的一个，广度优先遍历，暴力循环，和点赞最多的解答差不多，有进步，哈哈
func orangesRotting(grid [][]int) int {
	length := len(grid)
	if length == 0 {
		return -1
	}
	higth := len(grid[0])

	// 先找出所有坏橘子和好橘子
	badOrange := make([]*orange, 0)
	goodOrange := 0
	for i := 0; i < length; i++ {
		for j := 0; j < higth; j++ {
			if grid[i][j] == 2 {
				badOrange = append(badOrange, &orange{x: i, y: j})
			}
			if grid[i][j] == 1 {
				goodOrange++
			}
		}
	}

	// 一开始就没好的
	if goodOrange == 0 {
		return 0
	}

	// 开始模拟
	mark := 0
	for mark < math.MaxInt64 {
		// 第 mark 轮, mark初始化为0，但从1开始
		mark++

		size := len(badOrange)
		for _, v := range badOrange {
			// 看下上下左右，有没有好的橘子，如果又，感染它
			if v.x-1 >= 0 && grid[v.x-1][v.y] == 1 {
				badOrange = append(badOrange, &orange{x: v.x - 1, y: v.y})
				grid[v.x-1][v.y] = 2
				goodOrange--
			}
			if v.x+1 < length && grid[v.x+1][v.y] == 1 {
				badOrange = append(badOrange, &orange{x: v.x + 1, y: v.y})
				grid[v.x+1][v.y] = 2
				goodOrange--
			}
			if v.y-1 >= 0 && grid[v.x][v.y-1] == 1 {
				badOrange = append(badOrange, &orange{x: v.x, y: v.y - 1})
				grid[v.x][v.y-1] = 2
				goodOrange--
			}
			if v.y+1 < higth && grid[v.x][v.y+1] == 1 {
				badOrange = append(badOrange, &orange{x: v.x, y: v.y + 1})
				grid[v.x][v.y+1] = 2
				goodOrange--
			}
		}
		if size == len(badOrange) {
			// 这一轮没变化，感染结束

			// 有没有幸存的
			if goodOrange > 0 {
				return -1
			}

			mark--
			break
		}
		if goodOrange < 1 {
			// 没了
			break
		}

	}

	return mark
}

func Test_orangesRotting(t *testing.T) {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
	fmt.Println(orangesRotting([][]int{{0, 2}}))
	fmt.Println(orangesRotting([][]int{{1, 2}}))
}
