package _12

import "math"

func minSubArrayLen(target int, nums []int) int {
	left, right, windowSum := 0, 0, 0

	res := math.MaxInt32

	for right < len(nums) {
		windowSum += nums[right]
		for windowSum >= target {
			if res > right-left+1 {
				res = right - left + 1
			}
			windowSum -= nums[left]
			left++
		}

		right++
	}

	if res == math.MaxInt32 {
		return 0
	}

	return res
}
