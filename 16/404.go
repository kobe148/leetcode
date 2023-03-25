package _16

import "container/list"

func sumOfLeftLeaves(root *TreeNode) int {
	var sum = 0

	var preOrder func(*TreeNode, *TreeNode)
	preOrder = func(node *TreeNode, parent *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil && parent.Left == node {
			sum += node.Val
		}
		preOrder(node.Left, node)
		preOrder(node.Right, node)
	}

	preOrder(root, root)
	return sum
}

// BFS
func sumOfLeftLeaves2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var sum = 0
	// 用双向链表实现单向队列的功能
	var queue = list.New()
	queue.PushBack(root) // 入队(往链表表尾添加元素)
	for queue.Len() > 0 {
		// 每轮循环遍历处理一层的节点
		var size = queue.Len()
		for i := 0; i < size; i++ {
			// 出队，删除并得到链表表头元素
			var curr = queue.Remove(queue.Front()).(*TreeNode)
			if curr.Left != nil {
				if isLeafNode(curr.Left) {
					sum += curr.Left.Val
				} else {
					queue.PushBack(curr.Left)
				}
			}
			if curr.Right != nil && !isLeafNode(curr.Right) {
				queue.PushBack(curr.Right)
			}
		}
	}

	return sum
}

func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
