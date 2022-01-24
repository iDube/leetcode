package lc_300

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 278. 第一个错误的版本

var isBadVersion func(version int) bool

func firstBadVersion(n int) int {

	i, j := 1, n

	for i <= j {
		m := (i + j) / 2
		if isBadVersion(i) {
			return i
		} else if isBadVersion(m) {
			j = m
		} else {
			i = m + 1
		}
	}

	return n
}

func Test_isBadVersion(t *testing.T) {

	n, bad := 123, 45

	isBadVersion = func(version int) bool {
		return version >= bad
	}

	assertions := assert.New(t)
	assertions.Equal(firstBadVersion(n), bad)
}
