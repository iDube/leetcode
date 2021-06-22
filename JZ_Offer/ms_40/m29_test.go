package ms_40

// 29. 顺时针打印矩阵

func spiralOrder(matrix [][]int) []int {

	//设定边界，用方位了
	xStart := 0
	xEnd := len(matrix) - 1
	if xEnd < 0 {
		return nil
	}
	yStart := 0
	yEnd := len(matrix[0]) - 1

	var cache []int
	for {

		// 从左到右
		for i := yStart; i <= yEnd; i++ {
			cache = append(cache, matrix[xStart][i])
		}
		xStart++
		if xStart > xEnd {
			break
		}

		// 从上到下
		for i := xStart; i <= xEnd; i++ {
			cache = append(cache, matrix[i][yEnd])
		}
		yEnd--
		if yStart > yEnd {
			break
		}

		// 从右向左
		for i := yEnd; i >= yStart; i-- {
			cache = append(cache, matrix[xEnd][i])
		}
		xEnd--
		if xStart > xEnd {
			break
		}

		// 从下到上
		for i := xEnd; i >= xStart; i-- {
			cache = append(cache, matrix[i][yStart])
		}
		yStart++
		if yStart > yEnd {
			break
		}

	}

	return cache

}
