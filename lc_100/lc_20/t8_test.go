package lc_20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 8. 字符串转换整数 (atoi)

func myAtoi(str string) int {
	length := len(str)
	if length == 0 {
		return 0
	}

	num := 0
	flag := true
	int32Max := 1<<31 - 1
	int32Min := -1 << 31
	numStart := false
	_ = numStart

	for i := 0; i < length; i++ {
		c := int(str[i])
		if c == ' ' {
			if !numStart {
				continue
			} else {
				break
			}
		}

		if c == '-' {
			if !numStart {
				flag = false
				numStart = true
				continue
			} else {
				break
			}
		}
		if c == '+' {
			if !numStart {
				flag = true
				numStart = true
				continue
			} else {
				break
			}
		}
		if c < '0' || c > '9' {
			break
		} else {
			numStart = true
		}

		num = num*10 + c - '0'

		if num > int32Max {
			break
		}
	}

	if (flag && num < 0) || (!flag && num > 0) {
		num *= -1
	}

	if num > int32Max {
		return int32Max
	}
	if num < int32Min {
		return int32Min
	}

	return num

}

func TestMyAtoi(t *testing.T) {

	assertions := assert.New(t)

	assertions.Equal(myAtoi("1-0"), 1)
	assertions.Equal(myAtoi("0-1"), 0)                          // num 一只加，会得到个负值
	assertions.Equal(myAtoi("9223372036854775808"), 2147483647) // num 一只加，会得到个负值
	assertions.Equal(myAtoi("   -42"), -42)
	assertions.Equal(myAtoi("4193 with words"), 4193)
	assertions.Equal(myAtoi("words and 987"), 0)

}
