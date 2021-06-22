package ms_20

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func hammingWeight(n int) int {
	count := 0
	// int 默认为 int64,所以，1移位64次
	for i := 0; i < 64; i++ {
		tag := 1 << i
		if n&tag == tag {
			count++
		}
	}
	return count
}

func Test_hammingWeight(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(hammingWeight(0), 0)
	assertions.Equal(hammingWeight(1), 1)
	assertions.Equal(hammingWeight(2), 1)
	assertions.Equal(hammingWeight(3), 2)
	assertions.Equal(hammingWeight(4), 1)
	assertions.Equal(hammingWeight(5), 2)
	assertions.Equal(hammingWeight(9), 2)
	assertions.Equal(hammingWeight(21), 3)
	assertions.Equal(hammingWeight(33333), 6)
}

func TestByteCompute(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%v & 1 : %v \n", i, i&1)
	}

	n := 9
	for i := 0; i < 63; i++ {
		tag := 1 << i
		fmt.Printf("9 & %v : %v \n", tag, n&tag)
	}

}
