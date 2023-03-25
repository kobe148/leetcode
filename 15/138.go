package _15

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var nodeMap = make(map[*Node]*Node)

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	var oldNode = head
	var newNode = &Node{Val: head.Val}

	// bug 修复
	nodeMap[oldNode] = newNode

	var newHead = newNode
	for oldNode != nil {
		newNode.Next = getCloneNode(oldNode.Next)
		newNode.Random = getCloneNode(oldNode.Random)

		oldNode = oldNode.Next
		newNode = newNode.Next
	}

	return newHead
}

func getCloneNode(node *Node) *Node {
	if node == nil {
		return nil
	}

	if nodeMap[node] == nil {
		nodeMap[node] = &Node{Val: node.Val}
	}

	return nodeMap[node]
}
