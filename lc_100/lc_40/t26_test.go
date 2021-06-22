package lc_40

import (
	"fmt"
	"testing"
)

func removeDuplicates(nums []int) int {

	// 移出去第i个
	remove := func(ns []int, i int) {
		for i < len(ns)-1 {
			ns[i] = ns[i+1]
			i++
		}
	}

	length := len(nums)
	if length == 0 {
		return 0
	}

	i := 1
	list := nums[0]
	for i < length {
		if nums[i] == list {
			remove(nums, i)
			length--
		} else {
			list = nums[i]
			i++
		}
	}

	return length
}

func Test_removeDuplicates(t *testing.T) {
	a := []int{1, 1, 2}
	l := removeDuplicates(a)
	fmt.Println(l, a[0:l])

	a = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	l = removeDuplicates(a)
	fmt.Println(l, a[0:l])
}
