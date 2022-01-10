package lc_40

import (
	"fmt"
	"testing"
)

//  36. 有效的数独

func isValidSudoku(board [][]byte) bool {

	var rows, columns, boxes [9]map[byte]bool
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		columns[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == 46 {
				continue
			}

			// 计算在第几个box里
			boxIndex := 0
			boxIndex = i/3*3 + j/3

			// 判断是否存在
			if rows[i][num] || columns[j][num] || boxes[boxIndex][num] {
				return false
			}

			rows[i][num] = true
			columns[j][num] = true
			boxes[boxIndex][num] = true

		}
	}

	return true

}

func Test_isValidSudoku(t *testing.T) {

	board := [][]string{
		{"8", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"}}

	/**
	["8","3",".",".","7",".",".",".","."],
	["6",".",".","1","9","5",".",".","."],
	[".","9","8",".",".",".",".","6","."],
	["8",".",".",".","6",".",".",".","3"],
	["4",".",".","8",".","3",".",".","1"],
	["7",".",".",".","2",".",".",".","6"],
	[".","6",".",".",".",".","2","8","."],
	[".",".",".","4","1","9",".",".","5"],
	[".",".",".",".","8",".",".","7","9"]
	*/

	b := [][]byte{}
	for i := range board {
		tp := []byte{}
		for j := range board[i] {
			temp := []byte(board[i][j])
			tp = append(tp, temp[0])
		}
		b = append(b, tp)
	}

	fmt.Println(board)
	fmt.Println(b)

	fmt.Println(isValidSudoku(b))

	// TODO

}
