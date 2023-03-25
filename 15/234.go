package _15

/**
 * Definition for singly-linked list.
 *
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	newHead := slow.Next
	slow.Next = nil

	left, right := head, reverseList(newHead)
	for right != nil {
		if right.Val != left.Val {
			return false
		}

		right = right.Next
		left = left.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
