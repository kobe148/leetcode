package _01

// MoveZeroes 双指针
func MoveZeroes(nums []int) {
	first, last := 0, 0
	for last < len(nums) {
		if nums[last] == 0 {
			last++
		} else {
			nums[first], nums[last] = nums[last], nums[first]
			first++
			last++
		}
	}
}
