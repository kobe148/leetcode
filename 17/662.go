package _17

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	node  *TreeNode
	seqNo int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var maxWidth = 0
	// 用双向链表实现单向队列的功能
	var queue = list.New()
	// 给每个节点标序号
	queue.PushBack(&Node{node: root, seqNo: 1}) // 入队(往链表表尾添加元素)
	for queue.Len() > 0 {
		size := queue.Len()
		startSeqNo, endSeqNo := 0, 0
		for i := 0; i < size; i++ {
			curr := queue.Remove(queue.Front()).(*Node)
			node, seqNo := curr.node, curr.seqNo
			if i == 0 {
				startSeqNo = seqNo
			}
			if i == size-1 {
				endSeqNo = seqNo
			}
			if node.Left != nil {
				queue.PushBack(&Node{
					node:  node.Left,
					seqNo: 2 * seqNo,
				})
			}
			if node.Right != nil {
				queue.PushBack(&Node{
					node:  node.Right,
					seqNo: 2*seqNo + 1,
				})
			}
		}

		if endSeqNo-startSeqNo+1 > maxWidth {
			maxWidth = endSeqNo - startSeqNo + 1
		}
	}

	return maxWidth
}
