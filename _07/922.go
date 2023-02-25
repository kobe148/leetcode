package _07

func sortArrayByParityII(nums []int) []int {
	// 双指针 i 指向偶数， j 指向奇数
	n := len(nums)
	i, j := 0, 1

	for i < n {
		// 如果当前偶数位置是奇数元素的话
		if nums[i]%2 == 1 {
			// 那么在奇数位置上找到一个偶数，与之交换
			for nums[j]%2 == 1 {
				j += 2
			}

			nums[i], nums[j] = nums[j], nums[i]
		}
		i += 2
	}

	return nums
}
