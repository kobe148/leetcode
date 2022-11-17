package _01

func FindDisappearedNumbers(nums []int) []int {
	res := make([]int, 0, len(nums))
	n := len(nums)
	for i := 0; i < n; i++ {
		index := (nums[i] - 1) % n
		nums[index] += n
	}
	for i := 0; i < n; i++ {
		if nums[i] <= n {
			res = append(res, i+1)
		}
	}
	return res
}
