package _19

import "container/list"

func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		size := queue.Len()
		tmp := make([]int, 0, size)
		for i := 0; i < size; i++ {
			curr := queue.Remove(queue.Front()).(*Node)
			tmp = append(tmp, curr.Val)
			for _, child := range curr.Children {
				if child != nil {
					queue.PushBack(child)
				}
			}
		}
		res = append(res, tmp)
	}

	return res
}
