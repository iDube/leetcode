package lc_800

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 773. 滑动谜题

// BFS

func slidingPuzzle(board [][]int) int {

	// 切片转数组，方便比较
	tag := [2][3]int{{1, 2, 3}, {4, 5, 0}}
	var bd [2][3]int
	copy(bd[0][:], board[0])
	copy(bd[1][:], board[1])

	if bd == tag {
		return 0
	}

	type Job struct {
		BD   [2][3]int
		step int
	}

	// 初始状态
	queue := make([]Job, 0, 10)
	exit := make(map[[2][3]int]bool)
	queue = append(queue, Job{BD: bd, step: 0})
	exit[bd] = true

	// 为了当题目变成 n*m 时变动小点，后面不再出现 2*3
	n := len(bd)
	m := len(bd[0])

	for len(queue) > 0 {

		job := queue[0]
		queue = queue[1:]

		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {

				if job.BD[i][j] == 0 {

					// 0 与 上下左右交换，衍生出4个
					if i-1 >= 0 {
						bdNext := job.BD
						bdNext[i-1][j], bdNext[i][j] = bdNext[i][j], bdNext[i-1][j]
						if bdNext == tag {
							return job.step + 1
						}
						if exit[bdNext] == false {
							exit[bdNext] = true
							queue = append(queue, Job{BD: bdNext, step: job.step + 1})
						}

					}
					if j-1 >= 0 {
						bdNext := job.BD
						bdNext[i][j-1], bdNext[i][j] = bdNext[i][j], bdNext[i][j-1]
						if bdNext == tag {
							return job.step + 1
						}
						if exit[bdNext] == false {
							exit[bdNext] = true
							queue = append(queue, Job{BD: bdNext, step: job.step + 1})
						}
					}
					if i+1 < n {
						bdNext := job.BD
						bdNext[i+1][j], bdNext[i][j] = bdNext[i][j], bdNext[i+1][j]
						if bdNext == tag {
							return job.step + 1
						}
						if exit[bdNext] == false {
							exit[bdNext] = true
							queue = append(queue, Job{BD: bdNext, step: job.step + 1})
						}
					}
					if j+1 < m {
						bdNext := job.BD
						bdNext[i][j+1], bdNext[i][j] = bdNext[i][j], bdNext[i][j+1]
						if bdNext == tag {
							return job.step + 1
						}
						if exit[bdNext] == false {
							exit[bdNext] = true
							queue = append(queue, Job{BD: bdNext, step: job.step + 1})
						}
					}

				}

			}
		}

	}

	return -1
}

func Test_slidingPuzzle(t *testing.T) {

	assertions := assert.New(t)
	assertions.Equal(slidingPuzzle([][]int{{1, 2, 3}, {4, 5, 0}}), 0)
	assertions.Equal(slidingPuzzle([][]int{{1, 2, 3}, {4, 0, 5}}), 1)
	assertions.Equal(slidingPuzzle([][]int{{1, 2, 3}, {5, 4, 0}}), -1)
	assertions.Equal(slidingPuzzle([][]int{{4, 1, 2}, {5, 0, 3}}), 5)
	assertions.Equal(slidingPuzzle([][]int{{3, 2, 4}, {1, 5, 0}}), 14)

}
