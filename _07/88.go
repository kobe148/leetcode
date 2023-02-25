package _07

func merge_88(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1 // nums1的最后一个元素
	j := n - 1 // nums2的最后一个元素
	k := m + n - 1
	// 循环从右往左依次比较
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums2[k] = nums2[j]
			j--
		}
		k--
	}
}

func mergeSort2(nums []int, lo, hi int, tmp []int) {
	if lo >= hi {
		return
	}

	mid := (hi-lo)/2 + lo

	mergeSort2(nums, lo, mid, tmp)
	mergeSort2(nums, mid+1, hi, tmp)

	merge2(nums, lo, mid, hi, tmp)
}

func merge2(nums []int, lo, mid, hi int, tmp []int) {
	for i := lo; i <= hi; i++ {
		tmp[i] = nums[i]
	}

	i, j := lo, mid+1

	for k := lo; k <= hi; k++ {
		if i == mid+1 {
			nums[k] = tmp[j]
			j++
		} else if j == hi+1 {
			nums[k] = tmp[i]
			i++
		} else if tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i++
		} else {
			nums[k] = tmp[j]
			j++
		}
	}
}
