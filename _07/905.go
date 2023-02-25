package _07

func sortArrayByParity(nums []int) []int {
	// 快排分区
	less, great := 0, len(nums)-1

	for less < great {
		if nums[less]%2 > nums[great]%2 {
			nums[less], nums[great] = nums[great], nums[less]
		}

		if nums[less]%2 == 0 {
			less++
		}
		
		if nums[great]%2 == 1 {
			great--
		}
	}

	return nums
}
