package main

// swapPairs:
// head->[1]->[2]->[3]->[4]->[5]->[6]
// head->[2]->[1]->[4]->[3]->[6]->[5]
func swapPairs(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	newHead := head.Next

	next := head
	ptr := next.Next

	for {
		tempNext := ptr.Next
		ptr.Next = next
		if tempNext == nil {
			next.Next = nil
			break
		}
		if tempNext.Next == nil {
			next.Next = tempNext
			break
		}

		next.Next = tempNext.Next
		next = tempNext
		ptr = next.Next
	}

	return newHead
}
