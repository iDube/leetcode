package lc_6000

// 5412. 在既定时间做作业的学生人数

func busyStudent(startTime []int, endTime []int, queryTime int) int {

	// 题目给出，startTime 和 endTime 长度相等切都大于等于1

	count := 0

	for i := range startTime {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			count++
		}
	}

	return count

}
