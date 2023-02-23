package _06

func hammingDistance(x int, y int) int {
	// 异或求1的个数
	n := x ^ y

	count := 0

	for n != 0 {
		n = n & (n - 1) // 这里是 &，去掉最后一个 1
		count++
	}

	return count
}
