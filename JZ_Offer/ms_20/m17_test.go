package ms_20

import (
	"fmt"
	"math"
	"testing"
)

// 面试题17. 打印从1到最大的n位数

func printNumbers(n int) []int {
	max := int(math.Pow(10, float64(n))) - 1
	r := make([]int, max)

	for i := 0; i < max; i++ {
		r[i] = i + 1
	}
	return r
}

func Test_printNumbers(t *testing.T) {
	fmt.Println(printNumbers(0))
	fmt.Println(printNumbers(1))
	fmt.Println(printNumbers(2))
	fmt.Println(printNumbers(3))
}
