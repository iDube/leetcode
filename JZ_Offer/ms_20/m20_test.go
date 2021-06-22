package ms_20

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

/**
sss
*/

// 剑指 Offer 20. 表示数值的字符串

// 修修补补的解法，日了狗了，貌似也不对
func isNumber(s string) bool {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return false
	}
	// 只有一位时，必须时数字
	if len(s) == 1 {
		return s[0] >= '0' && s[0] <= '9'
	}

	if s[len(s)-1] == 'e' || s[len(s)-1] == '+' || s[len(s)-1] == '-' {
		return false
	}

	if s[0] < '0' || s[0] > '9' {
		if s[0] == '+' || s[0] == '-' {
			// 去符号
			s = s[1:]
		} else if s[0] != '.' {
			return false
		}
	}

	hasE := false
	hasDot := false
	hasLab := false

	for i := 0; i < len(s); i++ {

		if s[i] >= '0' && s[i] <= '9' {
			// 数字
			continue
		}

		// 只有一个e, 且 e前面是同一个数字
		if s[i] == 'e' {
			if hasE == true {
				return false
			}
			if i < 1 || !isNumber(s[:i]) {
				return false
			}

			hasE = true
			continue
		}

		// 只有一个. ，且在 e 前
		if s[i] == '.' {
			if hasE == true || hasDot == true {
				return false
			}
			hasDot = true
			continue
		}

		if s[i] == '+' || s[i] == '-' {

			if hasLab == true {
				return false
			}

			if i-1 < 0 || s[i-1] != 'e' {
				return false
			}

			hasLab = true
			continue
		}

		return false

	}

	return true
}

func TestIsNumber(t *testing.T) {
	assertions := assert.New(t)

	assertions.Equal(isNumber("+100"), true)
	assertions.Equal(isNumber("3.1415"), true)
	assertions.Equal(isNumber(" 100"), true)
	assertions.Equal(isNumber("5e2"), true)
	assertions.Equal(isNumber("-123"), true)
	assertions.Equal(isNumber("0123 "), true)
	assertions.Equal(isNumber("12e"), false)
	assertions.Equal(isNumber("1a3.14"), false)
	assertions.Equal(isNumber("1.2.3"), false)
	assertions.Equal(isNumber("+-5"), false)
	assertions.Equal(isNumber("-1E-16"), false)
	assertions.Equal(isNumber("12e+5.4"), false)
	assertions.Equal(isNumber("12."), true)
	assertions.Equal(isNumber(".1"), true)
	assertions.Equal(isNumber("."), false)
	assertions.Equal(isNumber("e"), false)
	assertions.Equal(isNumber("+"), false)
	assertions.Equal(isNumber(".e1"), false)
	assertions.Equal(isNumber(".1e1"), true)
	assertions.Equal(isNumber("1.e1"), true)
	assertions.Equal(isNumber("4e+"), false)

}

// 有限状态自动机
// 输入的符号： 空格，s +-符, d 数字, e
// 难点在定义 9 总状态及状态转移图，即下面的 states
func isNumber2(s string) bool {

	states := []map[int32]int{
		{' ': 0, 's': 1, 'd': 2, '.': 4}, // 0. start with 'blank'
		{'d': 2, '.': 4},                 // 1. 'sign' before 'e'
		{'d': 2, '.': 3, 'e': 5, ' ': 8}, // 2. 'digit' before 'dot'
		{'d': 3, 'e': 5, ' ': 8},         // 3. 'digit' after 'dot'
		{'d': 3},                         // 4. 'digit' after 'dot'(‘blank’ before 'dot')
		{'s': 6, 'd': 7},                 // 5. 'e'
		{'d': 7},                         // 6. 'sign' after 'e'
		{'d': 7, ' ': 8},                 // 7. 'digit' after 'e'
		{' ': 8},                         // 8. end with 'blank'
	}
	p := 0 // start with state 0
	for _, c := range s {
		t := '?' // unknown
		if '0' <= c && c <= '9' {
			t = 'd' // digit
		} else if c == '+' || c == '-' {
			t = 's' // sign
		} else if c == '.' || c == 'e' || c == ' ' {
			t = c // dot, e, blank
		}

		if _, ok := states[p][t]; !ok {
			return false
		}

		p = states[p][t]
	}

	return p == 2 || p == 3 || p == 7 || p == 8
}

func TestIsNumber2(t *testing.T) {
	assertions := assert.New(t)

	assertions.Equal(isNumber2("+100"), true)
	assertions.Equal(isNumber2("3.1415"), true)
	assertions.Equal(isNumber2(" 100"), true)
	assertions.Equal(isNumber2("5e2"), true)
	assertions.Equal(isNumber2("-123"), true)
	assertions.Equal(isNumber2("0123 "), true)
	assertions.Equal(isNumber2("12e"), false)
	assertions.Equal(isNumber2("1a3.14"), false)
	assertions.Equal(isNumber2("1.2.3"), false)
	assertions.Equal(isNumber2("+-5"), false)
	assertions.Equal(isNumber2("-1E-16"), false)
	assertions.Equal(isNumber2("12e+5.4"), false)
	assertions.Equal(isNumber2("12."), true)
	assertions.Equal(isNumber2(".1"), true)
	assertions.Equal(isNumber2("."), false)
	assertions.Equal(isNumber2("e"), false)
	assertions.Equal(isNumber2("+"), false)
	assertions.Equal(isNumber2(".e1"), false)
	assertions.Equal(isNumber2(".1e1"), true)
	assertions.Equal(isNumber2("1.e1"), true)
	assertions.Equal(isNumber2("4e+"), false)

}
