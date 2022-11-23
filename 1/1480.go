package _01

// RunningSum 前缀和
func RunningSum(nums []int) []int {
	res := make([]int, len(nums), len(nums))
	res[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		res[i] = nums[i] + res[i-1]
	}

	return res
}
