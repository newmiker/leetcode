package main

import (
	"github.com/stretchr/testify/assert"
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
