package lc_800

import (
	"fmt"
	"testing"
)

// 733. 图像渲染

type Point = [2]int

// 广度优先搜索 -> 队列
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	rSize := len(image)
	cSize := len(image[0])
	color := image[sr][sc]

	queue := []Point{}
	set := make(map[Point]bool)

	p := Point{sr, sc}
	queue = append(queue, p)
	set[p] = true

	for len(queue) > 0 {
		p = queue[0]
		queue = queue[1:]
		image[p[0]][p[1]] = newColor

		// p 的上方元素
		if p[0]-1 >= 0 && image[p[0]-1][p[1]] == color {
			pTemp := Point{p[0] - 1, p[1]}
			if !set[pTemp] {
				set[pTemp] = true
				queue = append(queue, pTemp)
			}
		}

		// p 的下方元素
		if p[0]+1 < rSize && image[p[0]+1][p[1]] == color {
			pTemp := Point{p[0] + 1, p[1]}
			if !set[pTemp] {
				set[pTemp] = true
				queue = append(queue, pTemp)
			}
		}

		// p 的左边元素
		if p[1]-1 >= 0 && image[p[0]][p[1]-1] == color {
			pTemp := Point{p[0], p[1] - 1}
			if !set[pTemp] {
				set[pTemp] = true
				queue = append(queue, pTemp)
			}
		}

		// p 的右边元素
		if p[1]+1 < cSize && image[p[0]][p[1]+1] == color {
			pTemp := Point{p[0], p[1] + 1}
			if !set[pTemp] {
				set[pTemp] = true
				queue = append(queue, pTemp)
			}
		}

	}

	return image
}

func TestName(t *testing.T) {

	image := [][]int{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 1},
	}
	fmt.Println(floodFill(image, 1, 1, 1))

}
