package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 != l2 {
		return false
	}
	has := make([][][]int, l1)
	for i := 0; i < l1; i++ {
		has[i] = make([][]int, l1)
		for j := 0; j < l1; j++ {
			has[i][j] = make([]int, l1+1)
		}
	}
	var help func(p1, p2, n int) bool
	help = func(p1, p2, n int) (f bool) {
		if has[p1][p2][n] == 1 {
			return true
		}
		if has[p1][p2][n] == -1 {
			return false
		}
		defer func() {
			if f {
				has[p1][p2][n] = 1
			} else {
				has[p1][p2][n] = -1
			}
		}()
		flag := [26]byte{}
		for i := 0; i < n; i++ {
			flag[s1[p1+i]-'a']++
			flag[s2[p2+i]-'a']--
		}
		for i := 0; i < 26; i++ {
			if flag[i] != 0 {
				return false
			}
		}
		if n < 4 {
			return true
		}
		//fmt.Println(s1[p1:p1+n],s2[p2:p2+n])
		for i := 1; i < n; i++ {
			if (help(p1, p2, i) && help(p1+i, p2+i, n-i)) || (help(p1, p2+n-i, i) && help(p1+i, p2, n-i)) {
				return true
			}
		}
		return false
	}
	return help(0, 0, len(s1))
}

func main() {
	fmt.Println(isScramble("eebaacbcbcadaaedceaaacadccd", "eadcaacabaddaceacbceaabeccd"))
}
