package _11

import (
	"container/heap"
	"sort"
)

// 大顶堆

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
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

func lastStoneWeight(stones []int) int {
	n := len(stones)
	if n == 1 {
		return stones[0]
	}

	pq := &hp{stones}
	heap.Init(pq)

	for pq.Len() > 1 {
		x, y := pq.pop(), pq.pop()
		if x > y {
			pq.push(x - y)
		}
	}

	if pq.Len() == 0 {
		return 0
	}

	return pq.IntSlice[0]
}
