package ms_80

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func strToInt(str string) int {

	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	// 符号
	sign := 1
	if str[0] > '9' || str[0] < '0' {
		if str[0] == '-' {
			sign = -1
		} else if str[0] == '+' {
			sign = 1
		} else {
			return 0
		}
	}

	// 缓存数字
	result := 0

	for i := range str {
		if i == 0 && (str[0] == '-' || str[0] == '+') {
			continue
		}
		if str[i] > '9' || str[i] < '0' {
			break
		}
		// temp = append(temp, int(str[i])-48)
		result = result*10 + sign*int((str[i])-'0')
		if result > 1<<31-1 {
			return 1<<31 - 1
		}
		if result < -1<<31 {
			return -1 << 31
		}

	}

	return result

}

func TestStrToInt(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(strToInt("  0000000000012345678"), 12345678)
	assertions.Equal(strToInt("20000000000000000000"), 2147483647)
	assertions.Equal(strToInt("42"), 42)
	assertions.Equal(strToInt("  -42"), -42)
	assertions.Equal(strToInt("4193 with words"), 4193)
	assertions.Equal(strToInt("  -91283472332"), -2147483648)
	assertions.Equal(strToInt("words and 987"), 0)
}
