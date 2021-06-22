package ms_40

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 剑指 Offer 31. 栈的压入、弹出序列

// 看来下解析，都是在模拟这个过程

func validateStackSequences(pushed []int, popped []int) bool {

	// 题目已经给出，popped 和 pushed 长度相同
	size := len(pushed)
	if size == 0 {
		return true
	}

	stack := make([]int, 0, size)
	i := 0

	// 一个个的入栈
	// 每次入栈后，都尝试安装出战序列进行出栈操作
	// 最后，如果栈为空，说明这个入栈序列和出栈序列是可以吻合的
	for _, num := range pushed {
		// 入栈
		stack = append(stack, num)

		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			// 循环判断出栈
			stack = stack[:len(stack)-1]
			i += 1
		}

	}

	return len(stack) == 0

}

func TestValidateStackSequences(t *testing.T) {
	assertions := assert.New(t)
	assertions.Equal(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}), true)
	assertions.Equal(validateStackSequences([]int{1, 5, 3, 4, 2}, []int{5, 4, 3, 2, 1}), true)
	assertions.Equal(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}), false)
}
