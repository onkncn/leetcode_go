package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	if l1 == 0 || l2 == 0 {
		return l2 + l1
	}
	mymin := func(a, b, c int) int {
		if a <= b && a <= c {
			return a
		} else if c <= a && c <= b {
			return c
		}
		return b
	}
	word1 = "a" + word1
	word2 = "a" + word2
	dp := make([][]int, l1+1)
	for i, _ := range dp {
		dp[i] = make([]int, l2+1)
	}
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if word1[i] == word2[j] {
				dp[i][j] = mymin(dp[i-1][j-1], dp[i-1][j]+1, dp[i][j-1]+1)
			} else {
				dp[i][j] = mymin(dp[i-1][j-1]+1, dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	fmt.Println(dp)
	return dp[l1][l2]
}

func main() {
	fmt.Println(minDistance("a", "aa"))
}
