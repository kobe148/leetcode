package _11

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	leftK, rightK := (m+n+1)/2, (m+n+2)/2
	lower := getKth(nums1, nums2, leftK)
	upper := getKth(nums1, nums2, rightK)

	return float64(lower+upper) * 0.5
}

func getKth(nums1 []int, nums2 []int, k int) int {
	var len1, len2, i, j = len(nums1), len(nums2), 0, 0

	for true {
		if i == len1 {
			return nums2[j+k-1]
		}
		if j == len2 {
			return nums1[i+k-1]
		}
		if k == 1 {
			return min(nums1[i], nums2[j])
		}

		// tips：计算 newi 和 newj 需要减 1 的原因：k/2 表示的是长度，长度是从 1 开始，而下标是从 0 开始的，所以需要减 1
		newi := min(i+(k/2), len1) - 1
		newj := min(j+(k/2), len2) - 1
		// tips：而下面计算 i 和 j 的时候加 1 的原因是：就是排除 i 前面或者 j 前面的元素，所以往前走一个
		// 计算 k 的时候加 1 的原因是：用从 0 开始的下标计算从 1 开始的长度，都需要加 1 的
		if nums1[newi] < nums2[newj] {
			k = k - (newi - i + 1)
			i = newi + 1
		} else {
			k = k - (newj - j + 1)
			j = newj + 1
		}
	}

	// 实际上这里是不会执行到，所以返回任何值都可以
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
