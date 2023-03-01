package _08

import "math"

func cutWood(k int, nums []int) int {
	// 找到最大值
	max := math.MinInt32
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	left, right := 1, max

	for left < right {
		// right-left变成right-left+1，否则会越界
		mid := left + (right-left+1)/2
		// mid长度可以截取的数量
		// 截取的数量大于k，证明截取的长度还不够长，所以需要尽可能更长一点
		if calWoodCnt(mid, nums) >= k {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}

func calWoodCnt(mid int, nums []int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		cnt += nums[i] / mid
	}

	return cnt
}
