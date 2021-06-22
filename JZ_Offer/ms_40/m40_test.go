package ms_40

import (
	"fmt"
	"testing"
)

// 40. 最小的k个数

// 一般就是堆排序喽
// 这题数据长度有限，不用堆也行呀,类插入排序喽

func getLeastNumbers(arr []int, k int) []int {

	if k <= 0 {
		return []int{}
	}
	if len(arr) <= k {
		return arr
	}

	rs := make([]int, k)
	for i := 0; i < k; i++ {
		rs[i] = 1<<63 - 1 // 初始化为最大数
	}

	insert := func(num int) {

		if rs[k-1] < num {
			return
		}

		rs[k-1] = num

		if k == 1 {
			return
		}

		for i := k - 1; i > 0; i-- {
			if rs[i] < rs[i-1] {
				rs[i], rs[i-1] = rs[i-1], rs[i]
			}
		}
	}

	for i := 0; i < len(arr); i++ {
		insert(arr[i])
	}

	return rs
}

func TestGetLeastNumbers(t *testing.T) {
	fmt.Println(getLeastNumbers([]int{3, 5, 6, 4, 3, 7, 8, 4, 23, 2, 5}, 4))
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))
	fmt.Println(getLeastNumbers([]int{0, 1, 2, 1}, 1))
	//[3,5,6,4,3,7,8,4,23,2,5]
	//4
}

// rs = {1，2，3，4，5}   k = 0
