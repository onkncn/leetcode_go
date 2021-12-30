package main

import "fmt"

func numDecodings(s string) int {
	l := len(s)
	if s[0] == '0' {
		return 0
	}
	if l == 1 {
		return 1
	}
	s = "0" + s
	dp := make([]int, l+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= l; i++ {
		if s[i-1] == '0' {
			if s[i] == '0' {
				return 0
			}
			dp[i] = dp[i-1]
		} else if s[i-1] == '1' {
			if s[i] == '0' {
				dp[i] = dp[i-2]
			} else {
				dp[i] = dp[i-1] + dp[i-2]
			}
		} else if s[i-1] == '2' {
			if s[i] == '0' {
				dp[i] = dp[i-2]
			} else if s[i] == '1' || s[i] == '2' || s[i] == '3' || s[i] == '4' || s[i] == '5' || s[i] == '6' {
				dp[i] = dp[i-1] + dp[i-2]
			} else {
				dp[i] = dp[i-1]
			}
		} else {
			if s[i] == '0' {
				return 0
			}
			dp[i] = dp[i-1]
		}
		//fmt.Println(s[1:i+1],dp[i])
	}
	return dp[l]
}

func main() {
	fmt.Println(numDecodings("110"))
}
