package _02

func ValidMountainArray(arr []int) bool {
	// 找到最大的那个数，并且不在第一个或者最后一个
	n, i := len(arr), 0
	for i < n-1 && arr[i] < arr[i+1] {
		i++
	}

	if i == 0 || i == n-1 {
		return false
	}

	for i < n-1 && arr[i] > arr[i+1] {
		i++
	}

	return i == n-1
}
