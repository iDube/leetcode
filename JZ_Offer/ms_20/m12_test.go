package ms_20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 面试题12. 矩阵中的路径

// 题目要求，返回 true or false, 这就比返回路径的数组下标简单了

// 方式1：暴力遍历 board， 找到 word[0], 记录位置，然后上下左右找word中其他字符。
//		 递归实现的话，将经过的路径改为一个word没有的字符，例如 $, 然后将 board 的位置和下一个word的位置传下去。
// 方式2：暴力遍历 board，记录所有word中的字符的位置。然后判断这些记录的位置能不能按 word 的顺序连成一条路径。

// 方式一：
func exist(board [][]byte, word string) bool {

	if len(word) == 0 {
		return false
	}

	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if dfs(board, word, x, y, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, x, y, i int) bool {
	// 矩阵边界检查
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return false
	}
	// 值检查
	if board[x][y] != word[i] {
		return false
	}
	// 最后一个字符
	if i+1 == len(word) {
		return true
	}

	// 用 / 占位，标记已经走过的路径
	temp := board[x][y]
	board[x][y] = '/'

	result := dfs(board, word, x-1, y, i+1) ||
		dfs(board, word, x+1, y, i+1) ||
		dfs(board, word, x, y-1, i+1) ||
		dfs(board, word, x, y+1, i+1)

	board[x][y] = temp

	return result

}

func Test_exist(t *testing.T) {
	assertions := assert.New(t)
	a := [][]byte{
		{'a', 'b', 'c', 'e'},
		{'s', 'f', 'c', 's'},
		{'a', 'd', 'e', 'e'}}

	assertions.Equal(exist(a, "abfc"), true)
	assertions.Equal(exist(a, "abfb"), false)
	assertions.Equal(exist(a, "bccs"), true)
	assertions.Equal(exist(a, "bcce"), true)
	assertions.Equal(exist(a, "eecs"), true)

}
