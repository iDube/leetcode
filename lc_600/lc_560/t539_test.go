package lc_560

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 539. 最小时间差

// 执行用时：280 ms, 在所有 Go 提交中击败了7.46%的用户
// 内存消耗：5.2 MB, 在所有 Go 提交中击败了38.81%的用户

func findMinDifference(timePoints []string) int {
	length := len(timePoints)

	// 转化为分
	m := make([]int, length)
	for i := range timePoints {
		temp := strings.Split(timePoints[i], ":")
		HH, _ := strconv.Atoi(temp[0])
		MM, _ := strconv.Atoi(temp[1])
		m[i] = HH*60 + MM
	}

	// 排序
	for i := 0; i < length; i++ {
		// 从 i 开始，每次循环一边，将剩余部分最小的移动到i
		for j := i + 1; j < length; j++ {
			if m[i] > m[j] {
				m[i], m[j] = m[j], m[i]
			}
		}
	}

	// 最小差值
	r := 1440 - (m[length-1] - m[0])
	for i := 0; i < length-1; i++ {
		temp := m[i+1] - m[i]
		if temp < r {
			r = temp
		}
	}

	return r
}

func Test_findMinDifference(t *testing.T) {

	assertions := assert.New(t)

	timePoints := [][]string{
		{"23:59", "00:00"},
		{"00:00", "23:59", "00:00"},
	}

	assertions.Equal(findMinDifference(timePoints[0]), 1)
	assertions.Equal(findMinDifference(timePoints[1]), 0)

}
