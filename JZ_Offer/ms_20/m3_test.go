package ms_20

import (
	"fmt"
	"testing"
)

// 面试题03. 数组中重复的数字

// 最先想到的是遍历，用map记录，找到一个就返回
func findRepeatNumber(nums []int) int {

	temp := make(map[int]int)
	for _, v := range nums {
		temp[v] = temp[v] + 1
		if temp[v] > 1 {
			return v
		}
	}
	return -1

}

func Test_findRepeatNumber(t *testing.T) {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
	fmt.Println(findRepeatNumber([]int{1, 0, 0, 6, 6, 1, 7, 2, 3, 1, 0, 2, 5, 3}))
	fmt.Println(findRepeatNumber([]int{1, 0, 0, 6, 7, 2, 3, 1, 0, 2, 5, 3}))
}
