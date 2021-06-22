package lc_6000

import (
	"fmt"
	"math"
	"testing"
)

// 5415. 圆形靶内的最大飞镖数量

// 数学问题吧
// 牛逼，圆心不确定，如果圆心是 (0,0),那么就是下面的解法了。
// 圆心不确定，就是 聚类问题了
// TODO 聚类

func numPoints(points [][]int, r int) int {
	count := 0
	for i := range points {
		length := math.Pow(float64(points[i][0]), 2) + math.Pow(float64(points[i][1]), 2)
		fmt.Println(length)
		if length <= math.Pow(float64(r), 2) {
			count++
		}
	}
	return count
}

func TestNumPoints(t *testing.T) {
	a := [][]int{{1, 2}, {3, 5}, {1, -1}, {2, 3}, {4, 1}, {1, 3}}
	fmt.Println(numPoints(a, 2))
}
