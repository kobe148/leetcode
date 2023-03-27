package _18

import "container/list"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 层序遍历
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			curr := queue.Remove(queue.Front()).(*Node)
			// 如果当前节点不是这一层的最后一个节点，则设置 next
			if i != size-1 {
				curr.Next = queue.Front().Value.(*Node)
			}
			if curr.Left != nil {
				queue.PushBack(curr.Left)
			}
			if curr.Right != nil {
				queue.PushBack(curr.Right)
			}
		}
	}

	return root
}
