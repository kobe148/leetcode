package _01

func RemoveDuplicates(nums []int) int {
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[fast] == nums[slow] {
			fast++
		} else {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1 // 因为slow是数组元素下标，所以数组总个数的话，是下标+1
}
