package _10

func largestRectangleArea(heights []int) int {
	var n = len(heights)
	// 1. 计算每根柱子左边第一个小于这根柱子的柱子(每根柱子的左边界)
	var left, stack = make([]int, n), []int{}

	for i := n - 1; i >= 0; i-- {
		left[i] = -1
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			left[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	// 2. 计算每根柱子右边第一个小于这根柱子的柱子(每根柱子的右边界)
	var right = make([]int, n)
	stack = []int{}
	for i := 0; i < n; i++ {
		right[i] = n
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	var ans = 0
	for mid := range heights {
		height := heights[mid]
		area := height * (right[mid] - left[mid] - 1)
		if area > ans {
			ans = area
		}
	}

	return ans
}
