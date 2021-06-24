package lc_180

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 149. 直线上最多的点数

func maxPoints(points [][]int) int {

	if len(points) < 3 {
		return len(points)
	}

	// 依次选取两点组成直线，判断其点对此直线的共线情况, 避免计算斜率，
	max := 2
	for i := 0; i < len(points); i++ {
		A := points[i]

		for j := i + 1; j < len(points); j++ {
			B := points[j]
			m := 2

			dx, dy := A[0]-B[0], A[1]-B[1]

			for k := j + 1; k < len(points); k++ {
				C := points[k]

				dx2, dy2 := A[0]-C[0], A[1]-C[1]
				if dx*dy2 == dx2*dy {
					m++
				}
			}

			if max < m {
				max = m
			}

		}

	}

	return max

}

func Test_maxPoints(t *testing.T) {

	tests := []struct {
		name string  // 测试名称
		args [][]int // 传入参数
		want int     // 预期数据
	}{
		{name: "没有点", args: [][]int{}, want: 0},
		{name: "只有一个点", args: [][]int{
			{1, 1},
		}, want: 1},
		{name: "只有两个点", args: [][]int{
			{1, 1},
			{0, 0},
		}, want: 2},
		{name: "有两个(1,1)的点", args: [][]int{
			{1, 1},
			{0, 0},
			{1, 1},
		}, want: 3},
		{name: "有两个点存在重复的情况(1,1), (0,0)", args: [][]int{
			{1, 1},
			{0, 0},
			{1, 1},
			{0, 0},
			{2, 3},
		}, want: 4},
		{name: "正常测试1", args: [][]int{
			{1, 1},
			{2, 2},
			{3, 3},
		}, want: 3},
		{name: "正常测试2", args: [][]int{
			{1, 1},
			{3, 2},
			{5, 3},
			{4, 1},
			{2, 3},
			{1, 4},
		}, want: 4},
		{name: "三个点相同的点", args: [][]int{
			{1, 1},
			{1, 1},
			{1, 1},
		}, want: 3},
		{name: "斜率为0", args: [][]int{
			{2, 3},
			{3, 3},
			{-5, 3},
			{2, 1},
		}, want: 3},
		{name: "斜率为无穷大", args: [][]int{
			{3, -1},
			{3, 2},
			{3, 1},
			{2, 1},
		}, want: 3},
		{name: "k除不尽", args: [][]int{
			{84, 250},
			{0, 0},
			{1, 0},
			{0, -70},
			{0, -70},
			{1, -1},
			{21, 10},
			{42, 90},
			{-42, -230},
		}, want: 6},
		{name: "浮点数精度丢失", args: [][]int{
			{94911152, 94911151},
			{0, 0},
			{94911151, 94911150},
		}, want: 2},
	}

	assertions := assert.New(t)

	for i, _ := range tests {
		assertions.Equal(maxPoints(tests[i].args), tests[i].want, tests[i].name)
	}

}
