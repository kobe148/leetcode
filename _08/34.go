package _08

func searchRange(nums []int, target int) []int {
	res := make([]int, 0, 2)
	left := getFirstEqualElement(nums, target)
	right := getLastEqualElement(nums, target)

	res = append(res, left)
	res = append(res, right)
	return res
}

func getFirstEqualElement(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] < target {
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

func getLastEqualElement(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			} else {
				left = mid + 1
			}
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
