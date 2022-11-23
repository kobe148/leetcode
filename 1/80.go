package _01

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	slow, fast := 2, 2
	for fast < len(nums) {
		if nums[fast] == nums[slow-2] {
			fast++
		} else {
			nums[slow] = nums[fast]
			slow++
			fast++
		}
	}

	return slow
}
