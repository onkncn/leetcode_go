package main

import "fmt"

func cal(s, p byte) bool {
	if s == p || p == '*' || p == '?' {
		return true
	}
	return false
}
func isMatch(s string, p string) bool {
	s = "!" + s
	p = "!" + p
	ls := len(s)
	lp := len(p)
	if ls == 0 {
		for i := 0; i < lp; i++ {
			if p[i] != '*' {
				return false
			}
		}
		return true
	}
	if lp == 0 {
		return false
	}
	dp := make([][]bool, ls)

	for i := 0; i < ls; i++ {
		dp[i] = make([]bool, lp)
	}
	dp[0][0] = cal(s[0], p[0])
	if p[0] == '*' {
		for i := 0; i < ls; i++ {
			dp[i][0] = true
		}
		for i := 1; i < lp; i++ {
			if p[i] == '*' {
				dp[0][i] = true
			} else {
				dp[0][i] = cal(s[0], p[i])
				break
			}
		}
	}
	if dp[0][0] {
		for i := 1; i < lp; i++ {
			if p[i] == '*' {
				dp[0][i] = true
			} else {
				break
			}
		}
	}
	for i := 1; i < ls; i++ {
		for j := 1; j < lp; j++ {
			if p[j] == '*' {
				dp[i][j] = dp[i-1][j-1] || dp[i][j-1] || dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-1] && cal(s[i], p[j])
			}
		}
	}
	return dp[ls-1][lp-1]
}
func main() {
	fmt.Println(isMatch("a", "a*"))
	fmt.Println(isMatch("a", "*a"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "m??*ss*?i*pi"))
	fmt.Println(isMatch("c", "*?*"))
}
