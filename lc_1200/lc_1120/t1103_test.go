package lc_1120

import (
	"fmt"
	"testing"
)

// 1103. 分糖果 II
// 分析：
// 1. 这是一个数列问题
// 2. 可一用暴力解法，就是不停的分，分完为止

// 这是我第一次写的，暴力解法，不停的分，分完为止
func distributeCandies(candies int, num_people int) []int {
	peopleCandies := make([]int, num_people)

	point := 0
	for candies > 0 {
		for i := 0; i < num_people; i++ {
			// 计算要分几个
			if candies <= point {
				point = candies
				candies = 0
			} else {
				point++
				candies -= point
			}

			// 分配
			peopleCandies[i] += point

			// 有没有分完
			if candies == 0 {
				break
			}
		}
	}
	return peopleCandies
}

func Test_distributeCandies(t *testing.T) {
	fmt.Println(distributeCandies(7, 4))
	fmt.Println(distributeCandies(10, 3))
	fmt.Println(distributeCandies(15, 3))
	fmt.Println(distributeCandies(90, 4))
}

// 然后我看了解题，发现暴力解法也有更优雅的实现方式
func distributeCandies2(candies int, num_people int) []int {

	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	// 这么多个小朋友
	peopleCandies := make([]int, num_people)
	i := 0
	for candies > 0 {
		// i%num_people 这点没想到，随着 i 的增长，i%num_people 一直是当前要发的小朋友位置的下标
		// min(i+1, candies) 就是当前要发的，因为i从0开始，琐事第i次要发的是i+1个，如果剩余的不足，就把剩余的全发了
		peopleCandies[i%num_people] += min(i+1, candies)

		// 减去这次发的
		candies -= min(i+1, candies)

		// 准备下一轮
		i++
	}

	return peopleCandies

}

func Test_distributeCandies2(t *testing.T) {
	fmt.Println(distributeCandies(7, 4))
	fmt.Println(distributeCandies(10, 3))
	fmt.Println(distributeCandies(15, 3))
	fmt.Println(distributeCandies(90, 4))

	fmt.Println(distributeCandies2(7, 4))
	fmt.Println(distributeCandies2(10, 3))
	fmt.Println(distributeCandies2(15, 3))
	fmt.Println(distributeCandies2(90, 4))
}
