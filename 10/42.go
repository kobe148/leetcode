package _10

func trap(height []int) int {
	stack := make([]int, 0)

	res := 0
	for i := 0; i < len(height); i++ {

		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			leftIndex := stack[len(stack)-1]
			currWidth := i - leftIndex - 1
			currHeight := min(height[leftIndex], height[i]) - height[top]
			res += currWidth * currHeight

		}

		stack = append(stack, i)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
