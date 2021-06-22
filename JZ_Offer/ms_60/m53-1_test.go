package ms_60

import "testing"

// 面试题53 - I. 在排序数组中查找数字 I

func search(nums []int, target int) int {
	temp := make(map[int]int)
	for i := range nums {
		temp[nums[i]]++
	}
	return temp[target]
}

// 有序数组，使用二分查找比较好，
// 左边界 left 和 右边界 right ，易得数字 target的数量为 right - left - 1
func search2(nums []int, target int) int {
	i, j := 0, len(nums)-1

	// 循环二分，找右边界 right
	for i <= j {
		m := (i + j) / 2
		if nums[m] <= target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	right := i
	println(right)

	// 没找到
	if j >= 0 && nums[j] != target {
		return 0
	}

	// 循环二分，找左边界 left
	i = 0
	for i <= j {
		m := (i + j) / 2
		if nums[m] < target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	left := j
	println(left)

	return right - left - 1
}

func TestSearch2(t *testing.T) {
	search2([]int{5, 7, 8, 8, 8, 9}, 8)
}
