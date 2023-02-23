package _06

func singleNumber_137(nums []int) int {
	once, twice := 0, 0
	for _, num := range nums {
		once = (once ^ num) & ^twice
		twice = (twice ^ num) & ^once
	}
	return once
}
