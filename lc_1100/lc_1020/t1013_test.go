package lc_1020

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 1013. 将数组分成和相等的三个部分

// 1. 求总和
// 2. 分段
// 3. 是否分尽
func canThreePartsEqualSum(A []int) bool {
	length := len(A)
	if len(A) < 3 {
		return false
	}

	sum := A[0]

	for i := 1; i < length; i++ {
		sum += A[i]
	}

	if sum%3 != 0 {
		return false
	}

	tagSum := sum / 3

	s := tagSum
	count := 0
	for i := 0; i < length; i++ {
		s -= A[i]
		if s == 0 {
			s = tagSum
			count++
		}
	}

	if s == tagSum && count == 3 {
		return true
	} else if s == tagSum && tagSum == 0 && count > 3 {
		return true
	} else {
		return false
	}

}

func Test_canThreePartsEqualSum(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(canThreePartsEqualSum([]int{1, 1, 1, 1}), false)
	assertions.Equal(canThreePartsEqualSum([]int{10, -10, 10, -10, 10, -10, 10, -10}), true)
	assertions.Equal(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}), true)
	assertions.Equal(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}), false)
	assertions.Equal(canThreePartsEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}), true)
}
