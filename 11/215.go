package _11

import (
	"container/heap"
	"sort"
)

// 小顶堆
type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] < h.IntSlice[j]
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func (h *hp) push(v int) {
	heap.Push(h, v)
}

func (h *hp) pop() int {
	return heap.Pop(h).(int)
}

func findKthLargest(nums []int, k int) int {
	pq := &hp{}
	for i := range nums {
		pq.push(nums[i])
		if pq.Len() > k {
			pq.pop()
		}
	}

	return pq.pop()
}
