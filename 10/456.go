package _10

import "math"

func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	stack := make([]int, 0, n)
	maxK := math.MinInt32

	for i := n - 2; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}

		for len(stack) > 0 && nums[i] < stack[len(stack)-1] {
			maxK = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		if nums[i] > maxK {
			stack = append(stack, nums[i])
		}
	}

	return false
}
