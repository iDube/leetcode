package lc_300

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 283. 移动零

// 1. 要保证非零元素的相对顺序，那么不能直接将零和末尾非零元素交换。
// 2. 像冒泡一样，将零慢慢移动到末尾
func moveZeroes(nums []int) {

	length := len(nums)

	// 记录0的个数，减少一定
	zoreCount := 0

	// i := length - 1 从后面找可以避免 0 和 0 交换
	for i := length - 1; i >= 0; i-- {
		if nums[i] != 0 {
			continue
		}

		// 找到零向后移动
		for j := i; j < length-1-zoreCount; j++ {
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}

		zoreCount++

	}

}

// 双指针
func moveZeroes2(nums []int) {

	i, j, length := 0, 0, len(nums)

	// 只要是非0的统统都赋给nums[j]，是0，i就向后移动
	for i < length {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
		i++
	}

	// j指针之后的都是0
	for j < length {
		nums[j] = 0
		j++
	}

}

func Test_moveZeroes(t *testing.T) {
	assertions := assert.New(t)

	n := []int{0, 1, 0, 3, 12}
	moveZeroes2(n)
	assertions.Equal(n, []int{1, 3, 12, 0, 0})

	n = []int{0}
	moveZeroes2(n)
	assertions.Equal(n, []int{0})

	n = []int{0, 0}
	moveZeroes2(n)
	assertions.Equal(n, []int{0, 0})

	n = []int{0, 0, 1, 3, 2, 6}
	moveZeroes2(n)
	assertions.Equal(n, []int{1, 3, 2, 6, 0, 0})
}
