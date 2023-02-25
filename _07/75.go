package _07

func sortColors(nums []int) {
	// 三路快排
	zero, two, i := 0, len(nums)-1, 0

	for i <= two {
		if nums[i] == 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			i++
			zero++
		} else if nums[i] == 2 {
			nums[i], nums[two] = nums[two], nums[i]
			two--
		} else {
			i++
		}
	}
}
