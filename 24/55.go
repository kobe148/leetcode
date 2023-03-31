package _24

func canJump(nums []int) bool {
	maxPos := 0
	for i := 0; i < len(nums); i++ {
		if maxPos < i {
			return false
		}
		if i+nums[i] > maxPos {
			maxPos = i + nums[i]
		}
	}

	return true
}
