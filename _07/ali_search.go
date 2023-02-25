package _07

import "math"

func search(nums []int) int {
	// 最大值，第二大值
	first, second := math.MinInt32, math.MinInt32
	for _, num := range nums {
		if num > first {
			second = first
			first = num
		} else if num < second {
			second = num
		}
	}

	return second
}
