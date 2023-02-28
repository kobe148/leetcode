package _08

func binarySearch(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	left, right := 0, len(nums)-1
	// 计算中间值

	for left <= right {
		// right + left 可能会溢出
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

/**
下面是有重复元素的二分查找
*/

// 查找第一个等于目标值的元素
func binarySearchForFirstTargetElement(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	// 计算中间值

	for left <= right {
		// right + left 可能会溢出
		mid := (right-left)/2 + left
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				right = mid - 1
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// 查找第一个大于等于目标值的元素
func binarySearchForFirstGETargetElement(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	// 计算中间值

	for left <= right {
		// right + left 可能会溢出
		mid := (right-left)/2 + left
		// 逻辑可以合并
		//if nums[mid] == target {
		//	if mid == 0 || nums[mid-1] < target {
		//		return mid
		//	} else {
		//		right = mid - 1
		//	}
		//} else if nums[mid] < target {
		//	left = mid + 1
		//} else { // target > nums[mid]
		//	if mid == 0 || nums[mid-1] < target {
		//		return mid
		//	}
		//	right = mid - 1
		//}
		// 逻辑合并之后的代码
		if target <= nums[mid] {
			if mid == 0 || nums[mid-1] < target {
				return mid
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}

	return -1
}

func binarySearchForLastLessTargetElement(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	// 计算中间值

	for left <= right {
		// right + left 可能会溢出
		mid := (right-left)/2 + left
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			} else {
				left = mid + 1
			}
		} else if target > nums[mid] {
			left = mid + 1
		} else { // target < nums[mid]
			right = mid - 1
		}
	}

	return -1
}

func binarySearchForLastTargetElement(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	// 计算中间值

	for left <= right {
		// right + left 可能会溢出
		mid := (right-left)/2 + left
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			} else {
				left = mid + 1
			}
		} else if target > nums[mid] {
			left = mid + 1
		} else { // target < nums[mid]
			right = mid - 1
		}
	}

	return -1
}

func binarySearchForLETargetElement(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	// 计算中间值

	for left <= right {
		// right + left 可能会溢出
		mid := (right-left)/2 + left
		if target >= nums[mid] {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			} else {
				left = mid + 1
			}
		} else { // target < nums[mid]
			right = mid - 1
		}
	}

	return -1
}
