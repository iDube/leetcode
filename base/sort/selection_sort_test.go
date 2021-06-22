package sort

import (
	"fmt"
	"testing"
)

func SelectionSort(arr []int) {
	length := len(arr)

	for i := 0; i < length; i++ {
		// 从 i 开始，每次循环一边，将剩余部分最小的移动到i
		for j := i + 1; j < length; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func TestSelectionSort(t *testing.T) {
	values := []int{1, 5, 4, 3, 6, 7, 7, 8, 4, 3, 2, 56, 8, 9}
	SelectionSort(values)
	fmt.Println(values)
}
