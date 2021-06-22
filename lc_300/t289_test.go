package lc_300

import "testing"

// 289. 生命游戏

// 因为同时发生，所以之前的状态要保留到后面的状态确定。

func gameOfLife(board [][]int) {

	x := len(board)
	if x == 0 {
		return
	}
	y := len(board[0])

	time2 := make([][]int, x)

	for i := 0; i < x; i++ {
		temp := make([]int, y)
		for j := 0; j < y; j++ {
			count := 0

			if i > 0 && j > 0 && board[i-1][j-1] > 0 {
				count++
			}
			if i > 0 && board[i-1][j] > 0 {
				count++
			}
			if i > 0 && j+1 < y && board[i-1][j+1] > 0 {
				count++
			}
			if j > 0 && board[i][j-1] > 0 {
				count++
			}
			if j+1 < y && board[i][j+1] > 0 {
				count++
			}
			if i+1 < x && j > 0 && board[i+1][j-1] > 0 {
				count++
			}
			if i+1 < x && board[i+1][j] > 0 {
				count++
			}
			if i+1 < x && j+1 < y && board[i+1][j+1] > 0 {
				count++
			}

			// 因为初始化为0，所以只管活的就行了
			if (count == 2 || count == 3) && board[i][j] == 1 {
				temp[j] = 1
			}
			if count == 3 && board[i][j] == 0 {
				temp[j] = 1
			}

		}
		time2[i] = temp
	}
	// board = time2
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			board[i][j] = time2[i][j]
		}
	}
}

func TestGameOfLife(t *testing.T) {
	a := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	gameOfLife(a)

}
