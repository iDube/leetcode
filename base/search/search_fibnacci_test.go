package search

import (
	"fmt"
	"testing"
)

func TestCreateFibnacci(t *testing.T) {
	fib := CreateFibnacci(20)
	fmt.Println(fib)
}

func TestSearchFibnacci(t *testing.T) {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18} //升序序列
	key := 10
	index := SearchFibnacci(slice, key)

	if index == -1 {
		fmt.Printf("%v不存在元素%v\n", slice, key)
	} else {
		fmt.Printf("%v位于%v下标为%v的位置。\n", key, slice, index)
	}

}

// 构建斐波那契数列
func CreateFibnacci(n int) []int {
	res := []int{1, 1}
	for i := 2; i < n; i++ {
		res = append(res, res[i-1]+res[i-2])
	}
	return res
}

// 斐波那契查找
func SearchFibnacci(slice []int, key int) int {
	n := len(slice)
	F := CreateFibnacci(20)

	// 1、计算n位于斐波那契数列的位置；k为斐波那契下标，需满足F(k)-1>=n
	k := 0
	for !(F[k]-1 >= n) {
		k++
	}

	// 2、将数组a扩展到F[k]-1的长度 构建新序列，多出位补slice[n-1]
	tempS := make([]int, F[k]-1)
	copy(tempS, slice)
	for i := n; i < len(tempS); i++ {
		tempS[i] = slice[n-1]
	}

	//3、开始斐波那契查找
	low, high := 0, n-1
	for low <= high {

		mid := low + F[k-1] - 1

		if tempS[mid] > key {
			high = mid - 1
			k -= 1 //查找值在前面的F(k-1)位中
		} else if tempS[mid] < key {
			low = mid + 1
			k -= 2 //查找值在后面的F(k-2)位中
		} else {
			if mid < n {
				return mid
			} else { //位于tempS的填补位
				return n - 1
			}
		}
	}
	return -1
}
