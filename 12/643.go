package _12

import "math"

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)

	ans := math.MinInt32
	left, right := 0, 0
	windowSum := 0

	for right < n {
		windowSum += nums[right]
		// 满足窗口的条件：达到了固定的窗口大小
		if right-left+1 == k {
			if windowSum > ans {
				ans = windowSum
			}
			windowSum -= nums[left]
			left++
		}

		right++
	}

	return float64(ans) / float64(k)
}
