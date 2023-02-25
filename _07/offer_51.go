package _07

func reversePairs(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	// 为了不改变原数组，这里先将原数组拷贝一份
	dest := make([]int, len(nums))
	copy(dest, nums)
	return reversePairsR(dest, 0, len(nums)-1, make([]int, len(nums)))
}

func reversePairsR(nums []int, lo int, hi int, tmp []int) int {
	if lo >= hi {
		return 0
	}
	mid := lo + (hi-lo)/2

	leftReversePairs := reversePairsR(nums, lo, mid, tmp)
	rightReversePairs := reversePairsR(nums, mid+1, hi, tmp)

	mergeReversePairs := mergeAndCountReversePairs(nums, lo, mid, hi, tmp)
	return leftReversePairs + rightReversePairs + mergeReversePairs
}

func mergeAndCountReversePairs(nums []int, lo, mid, hi int, tmp []int) int {
	for i := lo; i <= hi; i++ {
		tmp[i] = nums[i]
	}

	i, j := lo, mid+1
	count := 0

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

			count += mid - i + 1
		}
	}

	return count
}
