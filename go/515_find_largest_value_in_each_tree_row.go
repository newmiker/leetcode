package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	return walk(root, 0, make([]int, 0))
}

func walk(root *TreeNode, lvl int, memo []int) []int {
	if root == nil {
		return memo
	}
	if lvl > len(memo)-1 {
		memo = append(memo, root.Val)
	} else if root.Val > memo[lvl] {
		memo[lvl] = root.Val
	}

	memo = walk(root.Left, lvl+1, memo)
	memo = walk(root.Right, lvl+1, memo)

	return memo
}
