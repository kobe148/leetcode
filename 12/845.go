package _12

func longestMountain(arr []int) int {
	n, ans, left := len(arr), 0, 0

	for left+2 < n {
		right := left + 1
		// 先升后降
		if arr[left] < arr[right] {
			// 找最高点
			for right+1 < n && arr[right] < arr[right+1] {
				right++
			}

			if right+1 < n && arr[right] > arr[right+1] {
				// 找最低点
				for right+1 < n && arr[right] > arr[right+1] {
					right++
				}

				if right-left+1 > ans {
					ans = right - left + 1
				}
			} else {
				right++
			}

		}

		left = right
	}

	return ans
}
