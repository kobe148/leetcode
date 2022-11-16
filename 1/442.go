package _01

import "math"

// 输入：nums = [4,3,2,7,8,2,3,1]
// 输出：[2,3]
// 数组元素作为索引下标, 把元素作为下标，然后将对应下标的数据改为负数，这样下一次放到当前是负数时，证明之前已经被访问了，就是重复的
func FindDuplicates(nums []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		index := int(math.Abs(float64(nums[i]))) - 1
		if nums[index] < 0 {
			res = append(res, int(math.Abs(float64(nums[i]))))
		} else {
			nums[index] = -nums[index]
		}
	}

	return res
}

func FindDuplicates2(nums []int) []int {
	res := make([]int, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		index := (nums[i] - 1) % n // 对下标进行操作，防止越界，并且按照这个公式可以找到对应的元素
		nums[index] += n

	}

	for i := 0; i < n; i++ {
		if nums[i] > 2*n {
			res = append(res, i+1)
		}
	}

	return res
}
