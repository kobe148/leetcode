package _01

func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	maxSum := 0
	for left < right {
		sum := (right - left) * min(height[left], height[right])
		if sum > maxSum {
			maxSum = sum
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxSum
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
