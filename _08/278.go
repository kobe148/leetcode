package _08

func isBadVersion(version int) bool {
	return false
}

// 找到第一个等于
func firstBadVersion(n int) int {
	left, right := 1, n

	for left <= right {
		mid := left + (right-left)/2

		if isBadVersion(mid) {
			if mid == 1 || !isBadVersion(mid-1) { // nums[mid + 1] > target
				return mid
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}

	return 0
}
