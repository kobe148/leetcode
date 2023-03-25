package _16

import (
	"container/list"
	"math"
)

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res = make([]int, 0)
	// 用双向链表实现单向队列的功能
	var queue = list.New()
	queue.PushBack(root) // 入队(往链表表尾添加元素)
	for queue.Len() > 0 {
		// 每轮循环遍历处理一层的节点
		var size, maxValue = queue.Len(), math.MinInt32
		for i := 0; i < size; i++ {
			// 出队，删除并得到链表表头元素
			var curr = queue.Remove(queue.Front()).(*TreeNode)
			if curr.Val > maxValue {
				maxValue = curr.Val
			}
			if curr.Left != nil {
				queue.PushBack(curr.Left)
			}
			if curr.Right != nil {
				queue.PushBack(curr.Right)
			}
		}
		res = append(res, maxValue)
	}

	return res
}
