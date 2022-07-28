package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected []int
	}{
		{
			arr:      []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		}, {
			arr:      []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
	}
	for _, test := range tests {
		res := twoSum(test.arr, test.target)
		assert.Equal(t, test.expected, res)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		arr      []string
		expected string
	}{
		{
			arr:      []string{"flower", "flow", "flight"},
			expected: "fl",
		}, {
			arr:      []string{"dog", "racecar", "car"},
			expected: "",
		},
	}
	for _, test := range tests {
		res := longestCommonPrefix(test.arr)
		assert.Equal(t, test.expected, res)
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			s:        "a",
			t:        "ab",
			expected: false,
		},
	}
	for _, test := range tests {
		res := isAnagramOptimized(test.s, test.t)
		assert.Equal(t, test.expected, res)
	}
}

func TestSwapPairs(t *testing.T) {
	makeList := func(arr []int) *ListNode {
		node := &ListNode{Val: arr[0]}
		head := node

		for _, v := range arr[1:] {
			node.Next = &ListNode{Val: v}
			node = node.Next
		}

		return head
	}

	tests := []struct {
		list     *ListNode
		expected *ListNode
	}{
		{
			list:     makeList([]int{1}),
			expected: makeList([]int{1}),
		},
		{
			list:     makeList([]int{1, 2}),
			expected: makeList([]int{2, 1}),
		},
		{
			list:     makeList([]int{1, 2, 3, 4, 5, 6}),
			expected: makeList([]int{2, 1, 4, 3, 6, 5}),
		},
		{
			list:     makeList([]int{1, 2, 3, 4, 5}),
			expected: makeList([]int{2, 1, 4, 3, 5}),
		},
	}
	for _, test := range tests {
		resNode := swapPairs(test.list)
		expNode := test.expected

		for resNode != nil {
			require.Equal(t, resNode.Val, expNode.Val)

			resNode = resNode.Next
			expNode = expNode.Next
		}
	}
}
