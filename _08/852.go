package _08

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] < arr[mid+1] {
			// 上坡
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}
