package _25

func findLengthOfLCIS(nums []int) int {
	ans, slow := 1, 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast-1] >= nums[fast] {
			slow = fast
			continue
		}
		if fast-slow+1 > ans {
			ans = fast - slow + 1
		}
	}

	return ans
}
