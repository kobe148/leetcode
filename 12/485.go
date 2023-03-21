package _12

func findMaxConsecutiveOnes(nums []int) int {
	left, right := 0, 0
	res := 0

	for right < len(nums) {
		if nums[right] == 0 {
			if right-left > res {
				res = right - left
			}

			left = right + 1
		}

		right++
	}

	if right-left > res {
		res = right - left
	}

	return res
}
