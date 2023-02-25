package _07

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)

	cnt := [101]int{}

	for _, num := range nums {
		cnt[num]++
	}

	for i := 1; i < 101; i++ {
		cnt[i] += cnt[i-1]
	}

	ret := make([]int, n)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			ret[i] = 0
		} else {
			ret[i] = cnt[nums[i]-1]
		}
	}

	return ret
}
