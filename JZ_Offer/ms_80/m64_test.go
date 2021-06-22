package ms_80

import "math"

// 剑指 Offer 64. 求1+2+…+n

func sumNums(n int) int {
	return (int(math.Pow(float64(n), 2.0)) + n) >> 1
}
