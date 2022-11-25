package _02

func Candy(ratings []int) int {
	left, right := make([]int, len(ratings)), make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		left[i] = 1
		right[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if i != 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if i != len(ratings)-1 && ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}

	res := 0
	for i := 0; i < len(ratings); i++ {
		if left[i] > right[i] {
			res += left[i]
		} else {
			res += right[i]
		}

	}

	return res
}
