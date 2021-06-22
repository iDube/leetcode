package ms_20

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 面试题04. 二维数组中的查找

// 查找，有序数组，用二分查找吧；2纬，平面二分查找，哈哈
// 一开始总是从矩阵中间选数，想二分，没搞定，看了书，从右上角开始，逻辑简单多了；
// 如果没发套进去经典算法，先暴力点，完成题目更重要吧
// 且从右上角开始，感觉有点像二叉树，对比一次，排除一些
func findNumberIn2DArray(matrix [][]int, target int) bool {

	x := len(matrix)
	if x == 0 {
		return false
	}
	y := len(matrix[0])

	i := 0
	j := y - 1
	for i < x && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}

	return false

}

func Test_findNumberIn2DArray(t *testing.T) {
	ast := assert.New(t)

	a := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	ast.Equal(findNumberIn2DArray(a, 5), true)
	fmt.Println(findNumberIn2DArray(a, 11))
	fmt.Println(findNumberIn2DArray(a, 13))
	fmt.Println(findNumberIn2DArray(a, 29))
	fmt.Println(findNumberIn2DArray(a, 31))
	fmt.Println(findNumberIn2DArray(a, 1))

	b := [][]int{
		{1, 4, 7, 11, 15},
	}
	fmt.Println(findNumberIn2DArray(b, 9))
	fmt.Println(findNumberIn2DArray(b, 7))

	c := [][]int{
		{},
	}
	fmt.Println(findNumberIn2DArray(c, 9))
	fmt.Println(findNumberIn2DArray(c, 1))

	d := [][]int{}
	fmt.Println(findNumberIn2DArray(d, 1))

	e := [][]int{
		{-5},
	}
	ast.Equal(findNumberIn2DArray(e, -10), false)

}
