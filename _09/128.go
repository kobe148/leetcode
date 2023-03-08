package _09

// 解法一：排序
// 解法二：哈希
func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	ans := 1

	for _, num := range nums {
		
		if set[num-1] {
			continue
		}

		currNum, count := num, 1
		for set[currNum+1] {
			currNum += 1
			count += 1
		}

		if count > ans {
			ans = count
		}
	}

	return ans
}
