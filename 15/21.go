package _15

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	dummy := &ListNode{}

	curr := dummy

	for list1 != nil && list2 != nil {
		x, y := list1.Val, list2.Val
		if x <= y {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	if list1 == nil {
		curr.Next = list2
	}

	if list2 == nil {
		curr.Next = list1
	}

	return dummy.Next
}
