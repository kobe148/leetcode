package _11

import (
	"container/heap"
	"sort"
)

type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	k1 := KthLargest{k: k}
	for _, val := range nums {
		k1.Add(val)
	}

	return k1
}

func (this *KthLargest) Push(v interface{}) {
	this.IntSlice = append(this.IntSlice, v.(int))
}

func (this *KthLargest) Pop() interface{} {
	a := this.IntSlice
	v := a[len(a)-1]
	this.IntSlice = a[:len(a)-1]
	return v
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}

	return this.IntSlice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
