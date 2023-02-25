package _07

import "math"

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	// 1. 找到最大最小值
	max, min := nums[0], nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	if max == min {
		return 0
	}

	gap := int(math.Ceil(float64(max-min) / float64(len(nums)-1)))

	// 2. 初始化桶数组
	bucketNum := len(nums)
	buckets := make([][]int, bucketNum)
	for i := range nums {
		buckets[i] = make([]int, 2)
		buckets[i][0] = math.MaxInt32
		buckets[i][1] = math.MinInt32
	}

	// 3. 将所有元素添加到对应的桶中
	for _, num := range nums {
		// bucketId 计算逻辑如何理解，请参考 issue：https://gitee.com/douma_edu/douma_algo_training_camp/issues/I498BD
		bucketId := (num - min) / gap
		if num < buckets[bucketId][0] {
			buckets[bucketId][0] = num
		}
		if num > buckets[bucketId][1] {
			buckets[bucketId][1] = num
		}
	}

	// 4. 计算最大间隔
	maxGap, prevBucketsMax := 0, min
	for _, bucket := range buckets {
		if bucket[0] == math.MaxInt32 {
			continue
		}
		if bucket[0]-prevBucketsMax > maxGap {
			maxGap = bucket[0] - prevBucketsMax
		}
		prevBucketsMax = bucket[1]
	}

	return maxGap
}
