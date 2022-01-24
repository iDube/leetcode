package lc_1000

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 977. 有序数组的平方

// 排序
func sortedSquares(nums []int) []int {

	length := len(nums)

	for i := 0; i < length; i++ {
		nums[i] *= nums[i]
	}

	// 冒泡
	for i := length - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}

// 双指针，前提：题目已给出，nums 为有序数组，平方最大值只会出现在两端
func sortedSquares2(nums []int) []int {

	// 结果数组
	r := make([]int, len(nums))

	// i:左指针; j:右指针; wp: r写入指针
	i, j, wp := 0, len(nums)-1, len(nums)-1

	for i <= j {
		iq := nums[i] * nums[i]
		jq := nums[j] * nums[j]
		if iq < jq {
			r[wp] = jq
			j--
		} else {
			r[wp] = iq
			i++
		}
		wp--
	}

	return r

}

func Test_sortedSquares(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(sortedSquares2([]int{-4, -1, 0, 3, 10}), []int{0, 1, 9, 16, 100})
	assertions.Equal(sortedSquares2([]int{-7, -3, 2, 3, 11}), []int{4, 9, 9, 49, 121})
	assertions.Equal(sortedSquares2([]int{1}), []int{1})
	assertions.Equal(sortedSquares2([]int{}), []int{})
	assertions.Equal(sortedSquares2([]int{1, 2}), []int{1, 4})
	assertions.Equal(sortedSquares2([]int{2, 1}), []int{1, 4})
}
