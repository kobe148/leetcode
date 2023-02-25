package _07

func sortArray(nums []int) []int {
	// 归并排序、快速排序 ，大数据量考虑
	// sort.Ints(nums) 内置函数
	// 归并排序:拆解，合并，子问题
	mergeSort(nums, 0, len(nums)-1, make([]int, len(nums)))
	return nums
}

func mergeSort(nums []int, lo, hi int, tmp []int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2 // 防止溢出
	// 分别对左右进行递归排序处理
	mergeSort(nums, lo, mid, tmp)
	mergeSort(nums, mid+1, hi, tmp)

	// 排序完进行归并，合并处理
	merge(nums, lo, mid, hi, tmp)
}

func merge(nums []int, lo, mid, hi int, tmp []int) {
	// 要排序的数组拷贝到tmp数组中
	for i := lo; i <= hi; i++ {
		tmp[i] = nums[i]
	}

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		// 边界条件，左边遍历完，右边直接写入
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
