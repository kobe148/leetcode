package _15

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	l1 := lists[0]
	for i := 1; i < len(lists); i++ {
		l1 = mergeTwoLists(l1, lists[i])
	}

	return l1
}
