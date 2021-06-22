package sort

import (
	"fmt"
	"testing"
)

// 快速排序
// 参考：

func QuickSort(values []int, left int, right int) {

	if left < right {
		temp := values[left]

		i, j := left, right

		for {

			// 从右向左找，找到第一个比分水岭小的数
			for temp <= values[j] && i < j {
				j--
			}

			// 从左向右找，找到第一个比分水岭大的数
			for temp >= values[i] && i < j {
				i++
			}

			// 如果哨兵相遇，则退出循环
			if i >= j {
				break
			}

			// 交换左右两侧的值
			values[i], values[j] = values[j], values[i]

		}

		// 将分水岭移到哨兵相遇点
		values[left] = values[i]
		values[i] = temp

		// 递归，左右两侧分别排序
		QuickSort(values, left, i-1)
		QuickSort(values, i+1, right)
	}

}

func Test_QuickSort(t *testing.T) {
	values := []int{1, 5, 4, 3, 6, 7, 7, 8, 4, 3, 2, 56, 8, 9}
	QuickSort(values, 0, len(values)-1)
	fmt.Println(values)
}
