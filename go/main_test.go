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

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		roman    string
		expected int
	}{
		{roman: "III", expected: 3},
		{roman: "LVIII", expected: 58},
		{roman: "MCMXCIV", expected: 1994},
	}
	for _, test := range tests {
		res := romanToInt(test.roman)
		assert.Equal(t, test.expected, res)
	}
}

func TestMergeTwoSortedLists(t *testing.T) {
	tests := []struct {
		list1    *ListNode
		list2    *ListNode
		expected *ListNode
	}{
		{
			list1:    makeList([]int{1, 2, 4}),
			list2:    makeList([]int{1, 3, 4}),
			expected: makeList([]int{1, 1, 2, 3, 4, 4}),
		},
	}
	for _, test := range tests {
		resNode := mergeTwoLists(test.list1, test.list2)
		expNode := test.expected

		for resNode != nil {
			require.Equal(t, resNode.Val, expNode.Val)

			resNode = resNode.Next
			expNode = expNode.Next
		}
	}
}

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m        int
		n        int
		expected int
	}{
		{m: 2, n: 3, expected: 3},
		{m: 7, n: 3, expected: 28},
	}
	for _, test := range tests {
		res := uniquePaths(test.m, test.n)
		assert.Equal(t, test.expected, res)
	}
}

func TestLargestValues(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 9},
		},
	}
	expected := []int{1, 3, 9}
	res := largestValues(tree)
	assert.Equal(t, expected, res)
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums     []int
		elem     int
		expected int
	}{
		{nums: []int{1, 2, 2, 3}, elem: 5, expected: 4},
		{nums: []int{3, 2, 2, 3}, elem: 3, expected: 2},
		{nums: []int{}, elem: 3, expected: 0},
		{nums: []int{3}, elem: 3, expected: 0},
		{nums: []int{3, 3, 2, 3, 2}, elem: 3, expected: 2},
	}
	for _, test := range tests {
		k := removeElement(test.nums, test.elem)
		assert.Equal(t, test.expected, k)
	}
}
