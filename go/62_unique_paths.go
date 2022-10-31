package main

/*
There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 10^9.


Example 2:

Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down

_111111111
123456789
136х______
14х_______
15________

*/

func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}

	memo := make(map[[2]int]int, m*n)

	for i := 1; i < m; i++ {
		memo[[2]int{i, 0}] = 1
	}
	for i := 1; i < n; i++ {
		memo[[2]int{0, i}] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			memo[[2]int{i, j}] = memo[[2]int{i - 1, j}] + memo[[2]int{i, j - 1}]
		}
	}

	return memo[[2]int{m - 1, n - 1}]
}
