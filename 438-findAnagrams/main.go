package main

import "fmt"

func findAnagrams(s string, p string) []int {
	res := []int{}
	flag := [26]int{}
	ans := [26]int{}
	for _, c := range p {
		flag[c-'a']++
	}
	pl := len(p)
	sl := len(s)
	if sl < pl {
		return res
	}
	for i := 0; i < pl; i++ {
		c := s[i]
		flag[c-'a']--
	}
	if flag == ans {
		res = append(res, 0)
	}
	for i := pl; i < len(s); i++ {
		c := s[i]
		c_pre := s[i-pl]
		flag[c_pre-'a']++
		flag[c-'a']--
		if flag == ans {
			res = append(res, i-pl+1)
		}
	}
	return res
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
