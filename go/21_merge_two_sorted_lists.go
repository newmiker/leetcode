package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	ptr := head

	for {
		if list1 == nil {
			ptr.Next = list2
			break
		}
		if list2 == nil {
			ptr.Next = list1
			break
		}

		if list2.Val < list1.Val {
			ptr.Next = list2
			ptr = ptr.Next
			list2 = list2.Next
			continue
		}

		ptr.Next = list1
		ptr = ptr.Next
		list1 = list1.Next
	}

	return head.Next
}
