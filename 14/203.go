package _14

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// 虚拟节点
	dummy := &ListNode{Val: -1}
	dummy.Next = head

	curr := head
	prev := dummy

	for curr != nil {

		if curr.Val == val {
			prev.Next = curr.Next
			curr.Next = nil
		} else {
			prev = curr
		}

		// 往前移动
		curr = prev.Next
	}

	return dummy.Next
}
