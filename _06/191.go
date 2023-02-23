package _06

func hammingWeight(num uint32) int {
	count := uint32(0)

	for i := 0; i < 32; i++ {
		count += num & 1 // 跟每一位的1进行与，1 & 1 = 1

		num >>= 1
	}

	return int(count)
}
