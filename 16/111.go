package _16

import "container/list"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		level++
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return level
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}

	return level
}
