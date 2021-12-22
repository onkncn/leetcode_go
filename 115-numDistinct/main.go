package main

import "fmt"

func numDistinct(s string, t string) int {
	ls, lt := len(s), len(t)
	if lt == 0 {
		return 1
	}
	if ls == 0 {
		return 0
	}
	s = " " + s
	t = " " + t
	dp := make([][]int, ls+1)
	for i := 0; i <= ls; i++ {
		dp[i] = make([]int, lt+1)
	}
	for i := 0; i <= ls; i++ {
		for j := 0; j <= lt; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = 0
			} else if j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j]
				if s[i] == t[j] {
					dp[i][j] += dp[i-1][j-1]
				}
			}
		}
	}
	return dp[ls][lt]
}
func main() {
	fmt.Println((numDistinct("RRAA", "RA")))
}
