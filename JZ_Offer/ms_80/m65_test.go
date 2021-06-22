package ms_80

func add(a int, b int) int {

	// 假如 b是进位  a是非进位和
	if b == 0 {
		return a
	}

	// 进位赋值给c，准备下一次递归使用
	c := (a & b) << 1

	// 非进位和赋值给d ，准备下一次递归使用
	d := a ^ b

	return add(d, c)
}
