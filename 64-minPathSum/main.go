package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if i > 0 {
			dp[i] += dp[i-1]
		}
		dp[i] += grid[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				dp[j] = min(dp[j-1], dp[j])
			}
			dp[j] += grid[i][j]
		}
	}
	return dp[n-1]
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
