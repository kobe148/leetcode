package _02

func Rotate(nums []int, k int) {
	// 反转 k 次
	k = k % len(nums) // 这个很重要，如果k超过数组长度的话，取模跟K次旋转的效果是一样，也可以方式数组越界
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
