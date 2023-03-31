package _24

// 贪心策略：每步都选择能跳到的最远距离
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	step, start, end := 0, 0, 0
	for end < len(nums)-1 { // 走到最后一个位置的时候就不用走了
		maxPos := 0
		// 每次从 [start, end] 中都选择能跳到的最远距离
		for i := start; i <= end; i++ {
			if i+nums[i] > maxPos {
				maxPos = i + nums[i]
			}
		}
		start = end + 1
		end = maxPos
		step++
	}

	return step
}
