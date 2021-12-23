package main

import "fmt"

func fings(s string, length int) string {
	l := len(s)
	if length >= len(s) {
		return ""
	}
	mp := make(map[string]int)
	for i := length; i <= l; i++ {
		mp[s[i-length:i]]++
		if mp[s[i-length:i]] == 2 {
			return s[i-length : i]
		}
	}
	return ""
}
func longestDupSubstring(s string) string {
	l := len(s)
	h, t := 0, l
	res := ""
	for h < t {
		m := (h + t) / 2
		ans := fings(s, m)
		if ans != "" {
			res = ans
			fmt.Println(ans)
			h = m + 1
		} else {
			t = m
		}
	}
	return res
}

func main() {

}
