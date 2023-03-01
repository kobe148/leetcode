package _08

func findKthPositive(arr []int, k int) int {

	currNum, missCnt, lastMissNum, i := 1, 0, -1, 0

	for missCnt < k {
		if currNum == arr[i] { // 当等于最后一个元素的时候，不需要再i++，因为越界了
			if i+1 < len(arr) {
				i++
			}
		} else {
			missCnt++
			lastMissNum = currNum
		}
		currNum++
	}

	return lastMissNum
}
