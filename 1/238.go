package _01

func ProductExceptSelf(nums []int) []int {
	left, right := make([]int, len(nums), len(nums)), make([]int, len(nums), len(nums))
	res := make([]int, 0)
	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = nums[i-1] * left[i-1]
	}

	right[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}

	for i := 0; i < len(nums); i++ {
		res = append(res, left[i]*right[i])
	}
	return res
}
