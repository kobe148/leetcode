package _08

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[mid+1] { // 上坡
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}
