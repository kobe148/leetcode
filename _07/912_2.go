package _07

import "math/rand"

func sortArray_2(nums []int) []int {
	// 快速排序 ，大数据量考虑
	// sort.Ints(nums) 内置函数
	// 快速排序:分区，对分区的左右排序， 分区点的位置不变，分区点的两边是无序的，然后继续拆分成子问题

	quickSort(nums, 0, len(nums)-1)

	return nums
}

func quickSort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}

	// 分区点
	index := partition(nums, lo, hi) // 分区
	quickSort(nums, lo, index-1)
	quickSort(nums, index+1, hi)
}

// 找到分区点
func partition(nums []int, lo, hi int) int {
	i := rand.Intn(hi-lo+1) + lo // 随机选择一个点

	// 交换位置
	nums[hi], nums[i] = nums[i], nums[hi]

	pivot := nums[hi]

	less, great := lo, lo

	for ; great <= hi-1; great++ { // 保证less前面的元素都比pivot要小
		if nums[great] < pivot {
			nums[great], nums[less] = nums[less], nums[great]
			less++
		}
	}

	nums[hi], nums[less] = nums[less], nums[hi]

	return less
}

// 三路快拍
func partition_2(nums []int, lo, hi int) int {
	j := rand.Intn(hi-lo+1) + lo // 随机选择一个点

	// 交换位置
	nums[hi], nums[j] = nums[j], nums[hi]

	pivot := nums[hi]

	less, great := lo, hi

	i := lo

	for i <= great {
		if nums[i] < pivot {
			nums[less], nums[i] = nums[i], nums[less]
			less++
			i++
		} else if nums[i] > pivot {
			nums[great], nums[i] = nums[i], nums[great]
			great--
		} else {
			i++
		}
	}

	nums[hi], nums[less] = nums[less], nums[hi]

	return less
}
