package _02

func CheckPossibility(nums []int) bool {
	cnt := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			cnt++
			if cnt > 1 {
				return false
			}
			if i-2 >= 0 && nums[i] < nums[i-2] {
				// 比如：[4, 5, 2, 4] && i == 2
				nums[i] = nums[i-1]
			} else {
				// 比如：[-1,4,2,3]
				// 如果 i 指向的是 2，这个时候，nums[i] > nums[i - 2] 的，如果修改 nums[i] = nums[i - 1]，那么数组变为 [-1, 4, 4, 3]
				// 如果这个时候修改 nums[i - 1] = nums[i] 的话，那么数组变为 [-1, 2, 2, 3]，这样数组只需一次修改，就变为非递减数列了
				nums[i-1] = nums[i]
			}
		}
	}

	return true
}
