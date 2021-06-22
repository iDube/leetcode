package lc_6000

func luckyNumbers(matrix [][]int) []int {
	// 题目给出，x，y 都大于0
	x := len(matrix)
	if x == 0 {
		return []int{}
	}
	y := len(matrix[0])

	result := []int{}

	// 找出每行最小值
	// 记录列下标
	for i := 0; i < x; i++ {
		min := matrix[i][0]
		tagJ := 0
		for j := 1; j < y; j++ {
			// 题目给出，矩阵中每一个值都是不同的
			if min > matrix[i][j] {
				min = matrix[i][j]
				tagJ = j
			}
		}

		// 找出对应列最大值
		max := matrix[0][tagJ]
		for t := 1; t < x; t++ {
			if max < matrix[t][tagJ] {
				max = matrix[t][tagJ]
			}
		}

		if max == min {
			result = append(result, min)
		}

	}

	return result

}
