package _08

func search33(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if target == nums[mid] {
			return mid
		}

		// 确定目标值在哪个有序区间内，先确定哪个区间是有序
		if nums[left] <= nums[mid] {
			// 左边是有序区间
			// 上面已经判断了是否target == nums[mid]，所以这边不需要判断是否等于
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右边有序 nums[left] > nums[mid]
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
