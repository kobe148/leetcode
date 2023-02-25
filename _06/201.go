package _06

func rangeBitwiseAnd(left int, right int) int {
	shift := 0

	for left < right {
		left >>= 1
		right >>= 1
		shift += 1
	}

	return left << shift
}
