package _09

func firstMissingPositive(nums []int) int {
	var set = make(map[int]bool)
	for i := range nums {
		set[nums[i]] = true
	}

	for i := 1; i <= len(nums); i++ {
		if !set[i] {
			return i
		}
	}

	return len(nums) + 1
}
