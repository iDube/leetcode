package ms_60

import (
	"fmt"
	"testing"
)

// 面试题59 - I. 滑动窗口的最大值

// 感觉用大顶堆来存储划动窗口内的数据最合适,
// 这样划动时，执行移除左边的，插入右边的数据，然后就能获取下最大值了

// 或者用 59-2 内的可以获取最大值的队列

func maxSlidingWindow(nums []int, k int) []int {

	max := func(data []int) int {
		m := data[0]
		for i := 1; i < len(data); i++ {
			if m < data[i] {
				m = data[i]
			}
		}
		return m
	}

	// 边界值处理
	if k == 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}

	if k >= len(nums) {
		return []int{max(nums)}
	}

	result := make([]int, 0, len(nums)-k+1)

	i, j := 0, k
	for i < j && j <= len(nums) {
		result = append(result, max(nums[i:j]))
		i++
		j++
	}

	return result

}

func TestMaxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
