package _06

func getSum(a int, b int) int {
	// 异或:无进位的加法，与运算：算出进位
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}

	return a
}
