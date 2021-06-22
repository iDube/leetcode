package ms_80

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 面试题66. 构建乘积数组

func constructArr(a []int) []int {

	length := len(a)
	result := make([]int, 0, length)

	// 第一次循环，计算 a[i] 左侧部分乘积
	// i 为 0 时，左侧乘积定义为1
	for i, product := 0, 1; i < length; i++ {
		result = append(result, product)
		product *= a[i]
	}

	// 第二次循环，计算 a[i] 右侧部分乘积, 从右到左遍历
	// i 为 le 时，左侧乘积定义为1
	for i, product := length-1, 1; i >= 0; i-- {
		result[i] *= product
		product *= a[i]
	}

	return result
}

func TestConstructArr(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(constructArr([]int{1, 2, 3, 4, 5}), []int{120, 60, 40, 30, 24})
	assertions.Equal(constructArr([]int{}), []int{})
	assertions.Equal(constructArr([]int{1}), []int{1})
	assertions.Equal(constructArr([]int{2}), []int{1})
}
