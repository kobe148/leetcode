package _06

func singleNumber(nums []int) int {
	res := 0

	// 异或
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res
}
