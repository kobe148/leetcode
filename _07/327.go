package _07

func countRangeSum(nums []int, lower int, upper int) int {
	// 归并排序
	var prefixSum = make([]int, len(nums)+1)
	prefixSum[0] = 0
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	return mergeSort_327(prefixSum, 0, len(prefixSum)-1, make([]int, len(prefixSum)), lower, upper)
}

func mergeSort_327(prefixSum []int, lo int, hi int, tmp []int, lower int, upper int) int {
	if lo >= hi {
		return 0
	}
	mid := lo + (hi-lo)/2

	var leftSumCount = mergeSort_327(prefixSum, lo, mid, tmp, lower, upper)
	var rightSumCount = mergeSort_327(prefixSum, mid+1, hi, tmp, lower, upper)

	var count = 0
	// 计算当前有效的区间和个数
	i, l, r := lo, mid+1, mid+1
	for i <= mid {
		for l <= hi && prefixSum[l]-prefixSum[i] < lower {
			l++
		}
		for r <= hi && prefixSum[r]-prefixSum[i] <= upper {
			r++
		}
		count += (r - l) // 是r - l , 不是r - 1
		i++
	}

	merge_327(prefixSum, lo, mid, hi, tmp)
	return leftSumCount + rightSumCount + count
}

func merge_327(nums []int, lo int, mid int, hi int, tmp []int) {
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
