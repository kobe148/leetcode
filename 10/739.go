package _10

func dailyTemperatures(temperatures []int) []int {

	if len(temperatures) == 1 {
		return []int{0}
	}

	res := make([]int, len(temperatures))

	stack := make([]int, 0, len(temperatures))

	for i := range temperatures {

		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			res[top] = i - top
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}
