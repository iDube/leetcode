package ms_20

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// 面试题16. 数值的整数次方

// 0.00001, 2147483647 超出时间限制
func myPow1(x float64, n int) float64 {

	// 边界处理
	if n == 0 || x == 1 {
		return 1.0
	}

	flag := true
	if n < 0 {
		n = -n
		flag = false
	}
	result := x
	for i := 1; i < n; i++ {
		result *= x
	}

	if flag {
		return result
	} else {
		return 1 / result
	}

}

func myPow(x float64, n int) float64 {

	// 边界处理
	if n == 0 || x == 1 {
		return 1.0
	}

	flag := true
	if n < 0 {
		n = -n
		flag = false
	}
	result := 1.0
	for n > 0 {
		if n&1 == 1 {
			result *= x
		}
		x *= x
		n >>= 1
	}

	if flag {
		return result
	} else {
		return 1 / result
	}

}

func TestMyPow(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(myPow(2.0, 0), math.Pow(2.0, 0))
	assertions.Equal(myPow(1.0, 3), math.Pow(1.0, 3))
	assertions.Equal(myPow(2.0, 3), math.Pow(2.0, 3))
	assertions.Equal(myPow(2.3, 1500), math.Pow(2.3, 1500))
	assertions.Equal(myPow(0.00001, 2147483647), math.Pow(0.00001, 2147483647))
}
