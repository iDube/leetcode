package ms_20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO 面试题11. 旋转数组的最小数字

// 本质上是查找吧，但不是有序的，因为旋转了一下，凌乱了

// 看解析的答案
func minArray2(numbers []int) int {
	r := len(numbers) - 1
	l := 0
	for l < r {
		mid := (r + l) / 2
		if numbers[mid] > numbers[r] {
			l = mid + 1
		} else if numbers[mid] < numbers[r] {
			r = mid
		} else {
			// mid 等于 最 右边的值, 那就去掉右边的值，一开始没想到这点
			r--
		}
	}
	return numbers[l]
}

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		// 这种情况应该不存在
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	if len(numbers) == 2 {
		if numbers[0] >= numbers[1] {
			return numbers[1]
		} else {
			return numbers[0]
		}
	}

	l := len(numbers)
	mid := l / 2

	// 最早的做法，没想到处理等号
	/*if numbers[mid] >= numbers[l-1] {
		return minArray(numbers[mid:l])
	} else {
		return minArray(numbers[0 : mid+1])
	}*/

	if numbers[mid] > numbers[l-1] {
		return minArray(numbers[mid:l])
	} else if numbers[mid] < numbers[l-1] {
		return minArray(numbers[0 : mid+1])
	} else {
		// 中间和末尾是相等的，去掉末尾
		return minArray(numbers[0 : l-2])
	}
}

func Test_minArray(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(minArray([]int{3, 4, 5, 1, 2}), 1)
	assertions.Equal(minArray([]int{4, 5, 6, 7, 8, 0, 1, 2, 3}), 0)
	assertions.Equal(minArray([]int{7, 8, 0, 1, 2, 3, 4, 5, 6}), 0)
	assertions.Equal(minArray([]int{7}), 7)
	assertions.Equal(minArray([]int{7, 1}), 1)
	assertions.Equal(minArray([]int{6, 7, 1}), 1)
	assertions.Equal(minArray([]int{2, 2, 2, 2, 0, 1, 2}), 0)
	assertions.Equal(minArray([]int{2, 2, 2, 2, 2, 2, 2}), 2)
	assertions.Equal(minArray([]int{2, 2, 2, 2, 0, 0, 2}), 0)
	assertions.Equal(minArray([]int{1, 3, 3}), 1)
	assertions.Equal(minArray([]int{3, 1, 3, 3}), 1)
}
