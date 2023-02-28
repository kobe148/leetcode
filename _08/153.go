package _08

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[right] < nums[mid] {
			// 证明这边是一直在单调递增的不可能存在最小值
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
