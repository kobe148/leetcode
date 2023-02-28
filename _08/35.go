package _08

// 二分查找第一个大于等于 target 的索引
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			} else { // nums[mid-1] > target，需要向左边移动
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}

	return -1
}
