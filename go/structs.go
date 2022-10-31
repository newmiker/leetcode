package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(arr []int) *ListNode {
	node := &ListNode{Val: arr[0]}
	head := node

	for _, v := range arr[1:] {
		node.Next = &ListNode{Val: v}
		node = node.Next
	}

	return head
}
