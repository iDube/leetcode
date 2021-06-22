package lc_60

import (
	"errors"
	"fmt"
	"testing"
)

/**
TODO N 皇后问题
要点：深度优先搜索、递归、栈、全排列、位运算
分析：
思路：
	1. 找出的结果一定是一个全排列，这样这个全排列纵向上不冲突，因为每个数字分布在不同的行，横向上肯定不冲突。
	2. 验证第一步的全排列斜向上是否冲突。
	3.
*/

// TODO
// 栈，简单的实现
type Stack struct {
	data []int
}

func (s *Stack) String() string {
	// TODO
	return ""
}

// 出栈
func (s *Stack) pop() (int, error) {
	if len(s.data) < 1 {
		// 空了
		return 0, errors.New("stack is empty.")
	} else {
		point := len(s.data) - 1
		result := s.data[point]
		s.data = s.data[0:point]
		return result, nil
	}
}

// 入栈
func (s *Stack) push(d int) {
	s.data = append(s.data, d)
}

func (s *Stack) getAll() []int {
	return s.data
}

func (s *Stack) length() int {
	return len(s.data)
}

// 第一种解法
func solveNQueens(n int) [][]string {

	temp := &Stack{}
	fmt.Println(temp)

	// 验证
	verification := func(array []int, t int) bool {
		// TODO
		return false
	}

	for {
		for j := 0; j < n; j++ {
			// 验证
			if verification(temp.getAll(), j) {
				temp.push(j)
			} else {
				continue
			}
		}

		// 跑完以论，都不符合，???? 不能写这里吧
		if temp.length() < n {
			temp.pop()
		} else {
			break
		}
	}

	return nil
}

// 测试
func Test_solveNQueens(t *testing.T) {
	solveNQueens(3)
}

func Test_Stack(t *testing.T) {
	stack := &Stack{}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	fmt.Println(stack)
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}
